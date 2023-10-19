<script lang="ts">
    import { page } from "$app/stores";
    import { PUBLIC_API_URL } from "$env/static/public";
    import type { ProjectPageDTO } from "$lib/models/project";
    import { onMount } from "svelte";
    import Content from "./content.svelte";

    const projectID = $page.url.searchParams.get("id")
    let project:ProjectPageDTO
    let imgRef: string
    let dateTime: string

    onMount(async () => {
        let data:any = await fetch(`${PUBLIC_API_URL}/project?id=${projectID}`).then((res) => res.json())
        project = JSON.parse(data['message'])
        console.log(project)

        dateTime = new Date(project.Timestamp).toLocaleDateString("nl-NL")
        imgRef = project.BannerURL ? project.BannerURL : "404-error.png"
    })
</script>

{#if project && project.Content}
    <div class="mt-10 w-fit px-5 mx-auto max-w-4xl">
        <h1 class="h1 gradient-heading">{project.Title}</h1>
        <p class="italic text-end">{dateTime}</p>
        <img class="object-cover" src={imgRef} alt={project.BannerDescription}/>
        <div>
            <Content content={project.Content}/>
        </div>
    </div>
{/if}
