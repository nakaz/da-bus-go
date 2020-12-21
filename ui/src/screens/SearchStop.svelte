<script>
  import { navigate } from 'svelte-routing';
  let stop;

  const handleSearch = () => navigate(`/arrivals/${stop}`);
</script>

<div class='container'>
  <h1>Search Stop</h1>
  <div class='form-container'>
    <form
      on:submit={handleSearch}
      class='search-stop'
    >
      <input
        bind:value={stop}
        type='number'
        placeholder='Search Stop'
        autofocus
      >
      <button type='submit'>Search</button>
    </form>
  </div>
</div>

<style>
  .container {
    padding: 10px 10px;
  }
  .search-stop {
    display: flex;
    flex-direction: column;
  }
</style>
