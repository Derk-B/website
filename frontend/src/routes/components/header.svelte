<script lang="ts">
    import { AppBar, LightSwitch } from "@skeletonlabs/skeleton";
    import Terminal from '~icons/ic/outline-terminal'
</script>

<AppBar>
	<svelte:fragment slot="lead">
        <Terminal font-size="30"/>
    </svelte:fragment>
	<p class="text-2xl">Derk</p>
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