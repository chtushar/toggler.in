<script context="module" lang="ts">
import type { Load } from '@sveltejs/kit';
import axios from 'axios';

export const load: Load = async () => {
  await axios.get('http://localhost:9090/api/v1/users/user_info')
}

</script>
<h1>
  Dashboard
</h1>