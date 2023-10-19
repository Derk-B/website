import { PUBLIC_API_URL } from '$env/static/public'

export let fetchProjects = async function() {
    let data = await fetch(`${PUBLIC_API_URL}/projects`).then((apiData) => apiData.json())
    
    return JSON.parse(data['message'])
}