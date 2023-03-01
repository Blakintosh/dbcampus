<script lang="ts">
	
</script>

<div class="grid h-[100vh] w-full bg-slate-100 shadow-inner">
	<div class="place-self-center justify-self-center shadow-md p-4 bg-white rounded-md">
		<div class="flex flex-col items-center m-4">
			<img src="/pitstop_full.png" alt="logo" class="w-96" />
			<h2 class="text-xl font-medium p-1">Welcome back!</h2>
			<h3 class="text-xs text-slate-600">Please log in to your account to continue.</h3>
		</div>
		<div class="flex flex-col items-center m-4">
			<form>
				<div class="flex flex-col w-full">
					<label for="username" class="text-sm font-medium text-slate-600">Username</label>
					<input type="username" id="username" class="border border-slate-300 rounded-md p-2 mt-1 w-full" />
				</div>
				<div class="flex flex-col w-full mt-4">
					<label for="password" class="text-sm font-medium text-slate-600">Password</label>
					<input type="password" id="password" class="border border-slate-300 rounded-md p-2 mt-1 w-full" />
				</div>
				<div class="flex flex-col w-full mt-4">
					<button class="bg-slate-600 text-white font-medium text-sm rounded-md p-2 w-full hover:bg-slate-700 duration-75">Log in</button>
				</div>
			</form>
			<div class="flex flex-col items-center w-full mt-4">
				<a href="#" class="text-amber-500 font-medium text-sm underline">Haven't got an account? Register</a>
			</div>
		</div>
	</div>
</div>