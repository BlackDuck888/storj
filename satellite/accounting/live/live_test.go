// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package live

import (
	"context"
	"math/rand"
	"testing"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testrand"
	"storj.io/storj/satellite/accounting"
	"storj.io/storj/storage/redis/redisserver"
)

func TestPlainMemoryLiveAccounting(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	config := Config{
		StorageBackend: "plainmemory",
	}
	cache, err := NewCache(zaptest.NewLogger(t).Named("live-accounting"), config)
	require.NoError(t, err)

	// ensure we are using the expected underlying type
	_, ok := cache.(*plainMemoryLiveAccounting)
	assert.True(t, ok)

	projectIDs, sum, err := populateCache(ctx, cache)
	require.NoError(t, err)

	// make sure all of the "projects" got all space updates and got right totals
	for _, projID := range projectIDs {
		spaceUsed, err := cache.GetProjectStorageUsage(ctx, projID)
		require.NoError(t, err)
		assert.Equalf(t, sum, spaceUsed, "projectID %v", projID)
	}

	err = cache.ResetTotals(ctx)
	require.NoError(t, err)

	for _, projID := range projectIDs {
		spaceUsed, err := cache.GetProjectStorageUsage(ctx, projID)
		require.NoError(t, err)
		assert.EqualValues(t, 0, spaceUsed)
	}
}

func TestRedisLiveAccounting(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	address, cleanup, err := redisserver.Start()
	require.NoError(t, err)
	defer cleanup()

	config := Config{
		StorageBackend: "redis://" + address + "?db=0",
	}
	cache, err := NewCache(zaptest.NewLogger(t).Named("live-accounting"), config)
	require.NoError(t, err)

	// ensure we are using the expected underlying type
	_, ok := cache.(*redisLiveAccounting)
	assert.True(t, ok)

	projectIDs, sum, err := populateCache(ctx, cache)
	require.NoError(t, err)

	// make sure all of the "projects" got all space updates and got right totals
	for _, projID := range projectIDs {
		spaceUsed, err := cache.GetProjectStorageUsage(ctx, projID)
		require.NoError(t, err)
		assert.Equalf(t, sum, spaceUsed, "projectID %v", projID)
	}

	err = cache.ResetTotals(ctx)
	require.NoError(t, err)

	for _, projID := range projectIDs {
		spaceUsed, err := cache.GetProjectStorageUsage(ctx, projID)
		require.NoError(t, err)
		assert.EqualValues(t, 0, spaceUsed)
	}
}

func TestRedisCacheConcurrency(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	address, cleanup, err := redisserver.Start()
	require.NoError(t, err)
	defer cleanup()

	config := Config{
		StorageBackend: "redis://" + address + "?db=0",
	}
	cache, err := NewCache(zaptest.NewLogger(t).Named("live-accounting"), config)
	require.NoError(t, err)

	projectID := testrand.UUID()

	const (
		numConcurrent = 100
		inlineAmount  = 10
		remoteAmount  = 10
	)
	expectedSum := (inlineAmount * numConcurrent) + (remoteAmount * numConcurrent)

	var group errgroup.Group
	for i := 0; i < numConcurrent; i++ {
		group.Go(func() error {
			return cache.AddProjectStorageUsage(ctx, projectID, inlineAmount, remoteAmount)
		})
	}
	require.NoError(t, group.Wait())

	spaceUsed, err := cache.GetProjectStorageUsage(ctx, projectID)
	require.NoError(t, err)

	require.EqualValues(t, expectedSum, spaceUsed)
}

func TestPlainMemoryCacheConcurrency(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	config := Config{
		StorageBackend: "plainmemory",
	}
	cache, err := NewCache(zaptest.NewLogger(t).Named("live-accounting"), config)
	require.NoError(t, err)

	projectID := testrand.UUID()

	const (
		numConcurrent = 100
		inlineAmount  = 10
		remoteAmount  = 10
	)
	expectedSum := (inlineAmount * numConcurrent) + (remoteAmount * numConcurrent)

	var group errgroup.Group
	for i := 0; i < numConcurrent; i++ {
		group.Go(func() error {
			return cache.AddProjectStorageUsage(ctx, projectID, inlineAmount, remoteAmount)
		})
	}
	require.NoError(t, group.Wait())

	spaceUsed, err := cache.GetProjectStorageUsage(ctx, projectID)
	require.NoError(t, err)

	require.EqualValues(t, expectedSum, spaceUsed)
}

func TestResetTotals(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	config := Config{
		StorageBackend: "plainmemory:",
	}
	cache, err := NewCache(zaptest.NewLogger(t).Named("live-accounting"), config)
	require.NoError(t, err)

	// ensure we are using the expected underlying type
	_, ok := cache.(*plainMemoryLiveAccounting)
	assert.True(t, ok)

	projectID := testrand.UUID()
	err = cache.AddProjectStorageUsage(ctx, projectID, 0, -20)
	require.NoError(t, err)
}

func populateCache(ctx context.Context, cache accounting.LiveAccounting) (projectIDs []uuid.UUID, sum int64, _ error) {
	const (
		valuesListSize  = 1000
		valueMultiplier = 4096
		numProjects     = 200
	)
	// make a largish list of varying values
	someValues := make([]int64, valuesListSize)
	for i := range someValues {
		someValues[i] = int64((i + 1) * valueMultiplier)
		sum += someValues[i] * 2
	}

	// make up some project IDs
	projectIDs = make([]uuid.UUID, numProjects)
	for i := range projectIDs {
		projectIDs[i] = testrand.UUID()
	}

	// send lots of space used updates for all of these projects to the live
	// accounting store.
	errg, ctx := errgroup.WithContext(context.Background())
	for _, projID := range projectIDs {
		projID := projID
		errg.Go(func() error {
			// have each project sending the values in a different order
			myValues := make([]int64, valuesListSize)
			copy(myValues, someValues)
			rand.Shuffle(valuesListSize, func(v1, v2 int) {
				myValues[v1], myValues[v2] = myValues[v2], myValues[v1]
			})

			for _, val := range myValues {
				if err := cache.AddProjectStorageUsage(ctx, projID, val, val); err != nil {
					return err
				}
			}
			return nil
		})
	}

	return projectIDs, sum, errg.Wait()
}
