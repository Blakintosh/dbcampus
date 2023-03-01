<script lang="ts">
	import MainButton from "../../../components/common/ActionButton.svelte";
    import AuthenticationView from "../../../components/auth/AuthenticationView.svelte";
	import TextField from "../../../components/auth/TextField.svelte";
	import PasswordField from "../../../components/auth/PasswordField.svelte";

	let username: string = "";
	let password: string = "";
    let confirmPassword: string = "";

	let authenticating: boolean = false;

	const register = async () => {
		authenticating = true;
		const response = await fetch("/auth/registerPost", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			body: JSON.stringify({
				username: username,
				password: password,
			}),
		});

		const data = await response.json();
		if (data.status === 200) {
			alert(`Bro this is wild we got a successful response: ${JSON.stringify(data)}`);
		}
		authenticating = false;
	};
</script>


<svelte:head>
    <title>Register - Pitstop</title>
</svelte:head>

<AuthenticationView>
    <span slot="heading" class="w-full text-center">
        <h2 class="text-xl font-medium p-1">Registration</h2>
    </span>
    <span slot="body" class="w-full">
        <TextField label="Username" name="username" bind:value={username} />
        <PasswordField label="Password" name="password" bind:value={password} />
        <PasswordField label="Confirm Password" name="confirmPassword" bind:value={confirmPassword} />
        
        <div class="flex flex-col w-full mt-4">
            <MainButton label={authenticating ? "Registering..." : "Register"} on:click={register} loading={authenticating} />
        </div>
        <div class="flex flex-col items-center w-full mt-4">
            <a href="/auth/login" class="text-amber-500 font-medium text-sm underline">or go to Log in</a>
        </div>
    </span>
</AuthenticationView>