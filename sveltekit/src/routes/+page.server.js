import ApiClient from '../queries';

let albums = [
		{id: 0, title: "Default Album", artist: "Default Artist", price: 69.420},	
]
const apiClient = new ApiClient();

export async function load(){
	await apiClient.getAlbums.get()
      .then((res) => res ? res.json() : albums)
      .then((response) => { 
				console.log(response)
        albums = response.data
			})
			 .catch((err) => {
        console.error(err);
      });

		return {
			albums
		}
  }

// /** @type {import('./$types').Actions} */
// export const actions = {

// 	addAlbum: async ({ request }) => {
//     const formData = await request.formData();
// 		await apiClient.addAlbum.post(formData)
//     	.then((res) => res ? res.json() : false)
// 			.then(newAlbum => {
// 				if(newAlbum){
// 					albums = [...albums, newAlbum]
// 				}
// 			})
// 			.catch((err) => {
//     	  console.error(err);
//     	});
			
// 			return {
// 				success: true,
// 				formData,
// 				albums
// 			}
//     }
// };