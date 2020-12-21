<script>
import { Router, Link, Route } from 'svelte-routing';
  import { initClient } from '@urql/svelte';
  import Arrivals from './screens/Arrivals.svelte';
  import SearchStop from './screens/SearchStop.svelte';

  initClient({
    url: 'http://localhost:8080/graphql',
  });

</script>

<div>
  <Router>
    <Route path='/' component={SearchStop} />
    <Route path="/arrivals/:stop" let:params>
      <Arrivals stop="{params.stop}" />
    </Route>
  </Router>
</div>

