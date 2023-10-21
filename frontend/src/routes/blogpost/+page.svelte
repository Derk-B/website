<script lang="ts">
    import { page } from "$app/stores";
    import { PUBLIC_API_URL } from "$env/static/public";
    import type { BlogPageDTO } from "$lib/models/blog";
    import { onMount } from "svelte";
    import Content from "./content.svelte";

    const blogID = $page.url.searchParams.get("id")
    let blog:BlogPageDTO
    let imgRef: string
    let dateTime: string

    onMount(async () => {
        let data:any = await fetch(`${PUBLIC_API_URL}/blog?id=${blogID}`).then((res) => res.json())
        blog = JSON.parse(data['message'])
        console.log(blog)

        dateTime = new Date(blog.Timestamp * 1000).toLocaleDateString("nl-NL")
        if(!blog.BannerUrl) {
            imgRef = "404-error.png"
            return
        }

        imgRef = blog.BannerUrl.startsWith("http") ? blog.BannerUrl : `${PUBLIC_API_URL}/img/${blog.BannerUrl}`
    })
</script>

{#if blog && blog.Content}
    <div class="my-10 w-fit px-5 mx-auto max-w-4xl">
        <h1 class="h1 gradient-heading">{blog.Title}</h1>
        <p class="italic text-end">{dateTime}</p>
        <img class="object-cover m-auto" src={imgRef} alt={blog.BannerDescription}/>
        <div>
            <Content content={blog.Content}/>
        </div>
    </div>
{/if}
