<script lang="ts">
	import ApiClient from "../queries";
	import Card from '$lib/Card.svelte'

	const apiClient = new ApiClient();
	let albums = [
			{id: "0", title: "Default Album", artist: "Default Artist", price: 69.420},	
]
let newAlbum = {id: "0", title: "New Album", artist: "New Artist", price: 0}

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

	async function addAlbum(){
		await apiClient.addAlbum.post(newAlbum)
    .then((res) => res ? res.json() : false)
		.then(success => {
			if(success){
				albums = [...albums, newAlbum]
			}
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
			<Card>
				Album: {album.id}
				Title: {album.title}
				Artist: {album.artist}
				Price: {album.price}
			</Card>
		{/each}
	</div>

	<form>
		<label>
			<b>Title</b>
			<input name="title" type="text" bind:value={newAlbum.title}>
		</label>
		<label>
			<b>Artist </b>
			<input name="artist" type="text" bind:value={newAlbum.artist}>
		</label>
		<label>
			<b> Price </b>
			<input name="price" type="number" bind:value={newAlbum.price}>
		</label>
		<label>
			<button class="button" on:click={addAlbum}> Add Album </button>
		</label>
	</form>
</main>

<style>
	main {
		background-color: rgb(red, green, blue);
	}
	.album-list{
		display: flex;
	}

	/* Style inputs */
  form {
	display: flex;
	flex-direction: column;
  padding: 15px;
  margin: 8px 0;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box;
	max-width: 400px;
}
label{
	padding: 5px;
}
input{
	width: 350px;
}

</style>

