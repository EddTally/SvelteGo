<script lang="ts">
	import ApiClient from "../queries";
	import Card from '$lib/Card.svelte'

	const apiClient = new ApiClient();

	// This is the data that comes from the load function in +page.server.js
	export let data;

	// New album has no id, remember that
	let newAlbum = {title: "", artist: "", price: 0}
	let chosenAlbum = {id: "0", title: "", artist: "", price: 0}
	let getAlbumByIDUsed = false
	let albumMessage = "Album Found"
	let createAlbumMessage = ""
	let chosenAlbumID = 0

	async function addAlbum(){
		await apiClient.addAlbum
		.post(newAlbum)
    .then((res) => res ? res.json() : false)
		.then(res => {
			if(res){
				data.albums = [...data.albums, res.data]
				createAlbumMessage = res.message
			}
		})
		.catch((err) => {
      console.error(err);
    });
	}

	async function getAlbumByID(){
		await apiClient.getAlbumByID
		.get(chosenAlbumID)
    .then((res) => res?.json())
		.then(res => {
			console.log(res.data)
			chosenAlbum = res.data
			albumMessage = "Album Found"
			getAlbumByIDUsed = true
		})
		.catch((err) => {
      console.error(err);
			albumMessage = "Album Not Found"
			getAlbumByIDUsed = true
    });
	}
  

</script>

<main>
	<h1 style="title is-1">Welcome to SvelteKit</h1>
	<h2> Lets see if we can display our go albums here </h2>
	<!-- <button class="button" on:click={getAlbums}>Get Albums</button> -->
  	<!-- <button class="button" on:click={getDefault}>Get Default</button> -->
		<div class="album-list">
			{#each data.albums as album}
				<Card>
					<div class="album">
						<div class="album-keys">
							Album ID:  <br>
							Title:  <br>
							Artist:  <br>
							Price:  <br>
						</div>
						<div class="album-values">
							{album.id} <br>
							{album.title} <br>
							{album.artist} <br>
							{album.price} <br>
						</div>
					</div>
				</Card>
			{/each}
		</div>

	<div id="album-forms">
		<form>
			<label>
				<b>Title</b>
				<input 
				name="title" 
				type="text" 
				placeholder="New Title"
				bind:value={newAlbum.title}>
			</label>
			<label>
				<b>Artist </b>
				<input 
				name="artist" 
				type="text" 
				placeholder="New Artist"
				bind:value={newAlbum.artist}>
			</label>
			<label>
				<b> Price </b>
				<input 
				name="price" 
				type="number" 
				bind:value={newAlbum.price}
				>
			</label>
			<label>
				<button class="button" on:click={addAlbum}> Add Album </button>
			</label>
			{#if createAlbumMessage.length > 0}
				<p>{createAlbumMessage} </p>
			{/if}
		</form>

		<form>
			{#if getAlbumByIDUsed}
				<p> {albumMessage} </p>
			{/if}
			<label>
				<b> Select Album ID </b>
				<input 
				name="albumID" 
				type="number" 
				bind:value={chosenAlbumID}
				>
			</label>
			<label>
				<button class="button" on:click={getAlbumByID}> Get Album </button>
			</label>
				{#if getAlbumByIDUsed}
					<div class="chosenAlbum card">
						<div class="album">
							<div class="album-keys">
								Album ID:  <br>
								Title:  <br>
								Artist:  <br>
								Price:  <br>
							</div>
							<div class="album-values">
								{chosenAlbum.id} <br>
								{chosenAlbum.title} <br>
								{chosenAlbum.artist} <br>
								{chosenAlbum.price} <br>
							</div>
						</div>
					</div>
				{/if}
		</form>

</main>

<style>
	main {
		background-color: rgb(red, green, blue);
	}
	#album-forms{
		display: flex;
		flex-direction: row;
	}
	.album-list{
		display: flex;
		flex-wrap: wrap;
	}
	.album{
		display: flex;
		flex-direction: row;
	}
	.album-values{
		padding-left: 5px;
		margin-left: 2px;
	}
	.chosenAlbum{
		border: 1px red solid;
		padding: 2px
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

