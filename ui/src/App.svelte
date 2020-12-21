<script>
  import { initClient } from '@urql/svelte';
  import Arrivals from './screens/Arrivals.svelte';

  initClient({
    url: 'http://localhost:8080/graphql',
  });

</script>

<div>
  <Arrivals />
</div>

