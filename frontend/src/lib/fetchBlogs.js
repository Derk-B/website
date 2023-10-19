import { PUBLIC_API_URL } from '$env/static/public'

export let fetchBlogs = async function() {
    let data = await fetch(`${PUBLIC_API_URL}/blogposts`).then((apiData) => apiData.json())
    
    return JSON.parse(data['message'])
}