<script context="module" lang="ts">
	export const prerender = true;
</script>

<script lang="ts">
import SignIn from "$lib/SignIn.svelte";
import { goto } from "$app/navigation";

const redirectToDashboard = () => {
  goto('/dashboard')
}

</script>

<svelte:head>
	<title>Sign In | Toggler</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<SignIn on:success={redirectToDashboard} />

</section>
