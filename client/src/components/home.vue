<template>
	<div style="font-size: 18px; color: black; text-align: center; padding: 10%">
		You have successfully accessed this protected route <b>{{ user.Username }}</b> !!!
		<br>
		<button @click="logout">Logout</button>
	</div>
</template>
<script>
import store from "@/store.js"
import axios from "axios"
export default {
	data(){
		return{
			user: ""
		}
	},
	mounted(){
		axios({
			method: "get",
			url:"http://localhost:1000/auth/user",
			headers: { "X-Access-Token": store.state.token}
		})
		.then(res=> this.user = res.data)
		.catch(err=>alert(err.response.data || err.request))
	},
	methods: {
		logout(){
			store.state.token = ""
			this.$router.push("/")
		}
	}
};
</script>