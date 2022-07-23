<script lang="ts">
  import { createEventDispatcher } from 'svelte'
  import axios from 'axios';

  let email: string;
  let password: string;

  const dispatch = createEventDispatcher();

  const signin = async (event: any) => {
    event.preventDefault();
    try {
      const { data } = await axios.post('http://localhost:9090/api/v1/auth/signin', {
        email,
        password
      })

      if (data.data !== null) {
        dispatch('success')
      }
    } catch (error) {
      console.error(error)
    }
  }
</script>

<form on:submit={signin}>
  <input bind:value={email} type="email" placeholder="enter your email">
  <input bind:value={password} type="password" name="password" placeholder="enter your password">
  <button type="submit">Sign In</button>
</form>