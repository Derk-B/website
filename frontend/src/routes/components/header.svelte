<script lang="ts">
    import { AppBar, LightSwitch } from "@skeletonlabs/skeleton";
    import Terminal from '~icons/ic/outline-terminal'
	import MediaQuery from 'svelte-media-queries'
</script>

<MediaQuery query='(min-width: 700px)' let:matches>
	<AppBar gap=0>
		<svelte:fragment slot="lead">
			{#if matches}
			<Terminal font-size="30"/>
			{/if}
		</svelte:fragment>
		{#if matches}
			<p class="text-2xl">Derk</p>
		{/if}
		<svelte:fragment slot="trail">
			<nav>
				<ul class="flex flex-row gap-9 me-5">
					<li><a href="/">Home</a></li>
					<li><a href="/projects">Projects</a></li>
					<li><a href="/blog">Blog</a></li>
				</ul>
			</nav>
			<LightSwitch title="Switch color mode"/>
		</svelte:fragment>
	</AppBar>
</MediaQuery>