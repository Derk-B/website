<script>
	import '../app.postcss';
	import { AppShell } from '@skeletonlabs/skeleton';
	import Header from './components/header.svelte';
</script>

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header"><Header/></svelte:fragment>
	<svelte:fragment slot="sidebarLeft"></svelte:fragment>
	<svelte:fragment slot="pageHeader"></svelte:fragment>
	<slot />
	<svelte:fragment slot="pageFooter"></svelte:fragment>
</AppShell>
