<script lang="ts">
	import ApiClient from "../queries";

	const apiClient = new ApiClient();
	let albums = [
			{ID: "0", Title: "Default Album", Artist: "Default Artist", Price: 69.420},	
]

async function getAlbums() {
	await apiClient.getAlbums
      .get()
      .then((res) => {
				res ? res.json() : albums 
			})
      .then((data) => { 
				console.log(data)
			})
			 .catch((err) => {
        console.error(err);
      });
  }

</script>

<main>
	<h1 style="title is-1">Welcome to SvelteKit</h1>
	<h2> Lets see if we can display our go albums here </h2>
	<button class="button" on:click={getAlbums}>Get Albums</button>
	<div class="album-list">
		{#each albums as album}
			<div class="album-card">
				Album: {album.ID}
				Title: {album.Title}
				Artist: {album.Artist}
				Price: {album.Price}
			</div>
		{/each}
	</div>
</main>

<style>
	main {
		background-color: rgb(red, green, blue);
	}
	.album-list{
		display: flex;
	}
	.album-card{
		border: solid rgb(152, 151, 151) 2px;
		border-style: dashed;
		padding: 5px;
	}
</style>

