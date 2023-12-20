<script lang="ts">
    import { fetchBlogs } from "$lib/fetchBlogs";
    import type { BlogCardDTO } from "$lib/models/blog";
    import { onMount } from "svelte";
    import Error from "../components/error.svelte";
    import BlogCard from "./blogCard.svelte";
    
    let data : BlogCardDTO[];
    onMount(async () => {
        data = await fetchBlogs()
    })
</script>
{#if data}
<div class="grid md:grid-cols-2">
    {#each data as blog, _}
    <BlogCard blog={blog}/>
    {/each}
</div>
{/if}
