<script lang="ts">
	import ApiClient from "../queries";

	const apiClient = new ApiClient();
	let albums = [
			{id: "0", title: "Default Album", artist: "Default Artist", price: 69.420},	
]

async function getAlbums() {
	await apiClient.getAlbums.get()
      .then((res) => res ? res.json() : albums)
      .then((data) => { 
				console.log(data)
        albums = data
			})
			 .catch((err) => {
        console.error(err);
      });
  }
  
  async function getDefault() {
	await apiClient.getDefault
      .get()
      .then((res) => {
				console.log(res)
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
  	<button class="button" on:click={getDefault}>Get Default</button>
	<div class="album-list">
		{#each albums as album}
			<div class="album-card">
				Album: {album.id}
				Title: {album.title}
				Artist: {album.artist}
				Price: {album.price}
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

