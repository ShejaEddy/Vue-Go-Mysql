<template>
	<section id="login" v-bind:class="isShake">
 <form>
  <h2>Login</h2>
  <div class="info" v-bind:class="good">
   <p>{{ alert.message }}</p>
   <p v-show="user.username && user.password">{{ user.username }} / {{ user.password }}</p>
  </div>
  <input type="text" v-model="user.username" placeholder="Username" />
  <input type="password" v-model="user.password" placeholder="Password" />
  <button v-on:click="onSubmit">Log in</button>
 </form>
</section>
</template>
<script>
import axios from "axios"
import store from "@/store.js"
export default {
	data(){
	return {
		alert: {
			message: "Hello Folks!"
		},
		user: {
			username: "",
			password: ""
		},
		shake: false,
		good: ""
	}
		
	},
	computed: {
		isShake: function(){
			if(this.shake == true){
				return 'shake'
			}
			return 'none'
		}
	},
	methods: {
		onSubmit(e) {
			e.preventDefault();
			this.shake = false
			axios.post("http://localhost:1000/auth/login", this.user)
			.then(res => {
				this.good = "good"
				this.alert.message = "Success";
				store.state.token = res.data
				this.$router.push("/home")
			})
			.catch(err=>{
				this.shake = true;
				this.good = "error"
				if(err.response){
					this.alert.message = err.response.data;
					return
				}
				if(err.request) {
					console.log(err.request)
					return
				}
				console.log(err)
			})
				
		}
	}
};
	
</script>
<style>
	html, body, #app{
	width:100%;
	height:100%;
	margin:0px;
	font-family: 'Work Sans', sans-serif;
}

body{
		background-color: #f5f3f3;;
	  background-size: cover;
		display: flex;
		flex-direction:column;
		justify-content:center;
	  align-items:center;
		color: #fff;
}

section{
	margin-top: 10%;
	background-color: rgba(0, 0, 0, 0.72);
	width:25%;
	min-height:25%;
	display:flex;
	flex-direction:column;
	margin-left:auto;
	margin-right:auto;
}
form{
	display:flex;
	flex-direction:column;
	padding: 15px; 
}
h2{
	font-family: 'Archivo Black', sans-serif;
	color:#e0dada;
	margin-left:auto;
	margin-right:auto;
}

.info{
	padding:1em 5px;
	text-align:center;
	min-height:45px;
	display:flex;
	flex-direction:column;
	justify-content:center;
	align-items:center;
}

.info.error{
	background-color: #ff3c41;
}
.info p{
	margin:auto;
	padding:5px;
}
.info.good{
	border:1px solid #416d50;
	background-color: #47cf73;
	color:#416d50;
}

input{
	height:35px;
	padding: 5px 5px;
	margin: 10px 0px;
	background-color:#e0dada;
	border:none;
}
button{
	height:40px;
	padding: 5px 5px;
	margin: 10px 0px;
	font-weight:bold;
	background-color:#be5256;
	border:none;
	color:#e0dada;
	cursor:pointer;
	font-size:16px;
}
button:hover{
	background-color:#711f1b;
}

@-webkit-keyframes shake{
	from, to{
		-webkit-transform: translate3d(0, 0, 0);
		transform:translate3d(0,0,0);
	}
	10%,30%,50%,70%,90%{
		-webkit-transform: translate3d(-10px, 0, 0);
		transform:translate3d(-10px,0,0);
	}
	20%,40%,60%,80%{
		-webkit-transform: translate3d(10px, 0, 0);
		transform:translate(10px,0,0);
	}
}

.shake{
	animation-name: shake;
	animation-duration:1s;
	/*animation-fill-mode: both;*/
}

@media screen and (max-width: 780px) {
	section{
		width:90%;
	}
}
</style>