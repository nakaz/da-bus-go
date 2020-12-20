<script>
const list = new Array(10);
const mockData = list.fill({
  "headsign": "KALIHI TRANSIT CENTER",
  "latLng": [21.3023, -157.84484],
  "arrivalTime": "2020-12-20T11:50:00-10:00",
  "route": "1",
  "estimated": "1",
});
</script>

<div class='arrivals-list'>
  {#each mockData as arrival}
    <div class='arrivals-list-item'>
      <div class="arrivals-list-item--container">
        <div class='arrivals-list-item--top'>
          <span>{arrival.route}</span>
          -
          <span>{arrival.headsign}</span>
        </div>
        <div class='arrivals-list-item--bottom'>
          <span>{arrival.arrivalTime}</span>
        </div>
      </div>
      <div class="arrivals-list-item--eta">
        {arrival.estimated}
      </div>
    </div>
  {/each}
</div>

<style>
.arrivals-list {
  display: flex;
  flex-direction: column;
}
.arrivals-list-item {
  display: flex;
  border-bottom: 1px solid black;
  justify-content: space-between;
}
.arrivals-list-item--container {
  display: flex;
  padding: 10px;
  flex-direction: column;
}
.arrivals-list-item--eta {
  display: flex;
  padding: 10px;
  align-items: center;
  justify-content: center;
}
</style>
