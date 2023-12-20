<script lang="ts">
    import { PUBLIC_API_URL } from "$env/static/public";
    import type { BlogCardDTO } from "$lib/models/blog";
    import { onMount } from "svelte";

    export let blog:BlogCardDTO

    let dateTime:string
    let imgURL:string

    onMount(() => {
        let date:Date = new Date(blog.Timestamp * 1000)
        dateTime = date.toLocaleDateString("nl-NL")

        if(!blog.BannerUrl) {
            imgURL = "404-error.png"
            return
        }

        imgURL = blog.BannerUrl.startsWith("http") ? blog.BannerUrl : `${PUBLIC_API_URL}/img/${blog.BannerUrl}`
    })
</script>

<a href="/blogpost?id={blog.ID}" class="card p-4 mx-auto my-5 min-w-[40%] w-[90%]">
    <img src={imgURL} class="mx-auto max-h-96" alt={blog.BannerAlt}/>
    <h2 class="h2 mb-5">{blog.Title}</h2>
    <p>{blog.Description}</p>
    <p class="text-end">{dateTime}</p>
</a>