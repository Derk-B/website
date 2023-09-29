<script lang="ts">
  let links:any[] = [
    {
      id: 1,
      url: "/post",
      name: "My thoughts on UML data diagrams"
    }
  ];

  // Fetch data from the API
  async function fetchLinks() {
    try {
      const response = await fetch('https://your-api-url.com/links'); // Replace with your API URL
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      links = data; // Assuming the API response is an array of link objects
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  }

  // Fetch data when the component is mounted
  import { onMount } from 'svelte';
  onMount(fetchLinks);
</script>

<!-- Render the list of links -->
<ul>
  {#each links as link (link.id)}
    <li>
      <a href={link.url}>{link.name}</a>
    </li>
  {/each}
</ul>

