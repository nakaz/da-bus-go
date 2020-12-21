<script>
  import { operationStore, query } from '@urql/svelte';

  export let stop;

  const arrivals = operationStore(`
    query ($stop: Int!){
      arrivals(stop: $stop) {
        headsign
        latLng
        arrivalTime
        route
        estimated
        vehicleId
      }
    }
  `, { stop: Number(stop) });
  query(arrivals);
</script>

<div class='arrivals-header'>
  Stop Name / Street: Waialae Ave & Harding St.
  Stop Location: NA
  Stop ID: 12312 - W Bound
  Route: 1
</div>
<div class='arrivals-list'>
  {#if $arrivals.fetching}
    <p>Loading...</p>
  {:else if $arrivals.error}
    <p>{$arrivals.error.message}</p>
  {:else}
    {#each arrivals.data.arrivals as arrival}
      <div class='item'>
        <div class="container">
          <div>
            <span>{arrival.route}</span>
            -
            <span>{arrival.headsign}</span>
          </div>
          <div>
            <span>{arrival.arrivalTime}</span>
          </div>
        </div>
        <div class="eta">
          {arrival.estimated}
        </div>
      </div>
    {/each}
  {/if}
</div>

<style>
.arrivals-header {
  padding: 10px;

}
.arrivals-list {
  display: flex;
  flex-direction: column;

  & .item {
    display: flex;
    border-bottom: 1px solid black;
    justify-content: space-between;

    & .container {
      display: flex;
      padding: 10px;
      flex-direction: column;
    }

    & .eta {
      display: flex;
      padding: 10px;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
