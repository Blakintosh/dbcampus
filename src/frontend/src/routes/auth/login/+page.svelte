<script lang="ts">
	import { goto } from "$app/navigation";
	import AuthenticationView from "../../../components/auth/AuthenticationView.svelte";
	import PasswordField from "../../../components/auth/PasswordField.svelte";
	import TextField from "../../../components/auth/TextField.svelte";
    import MainButton from "../../../components/common/ActionButton.svelte";

	let username: string = "";
	let password: string = "";

    let usernameIsError = false;
    let usernameErrorMessage = "";

    let passwordIsError = false;
    let passwordErrorMessage = "";

	let authenticating: boolean = false;

    const resetAuthErrors = () => {
        usernameIsError = false;
        usernameErrorMessage = "";
        passwordIsError = false;
        passwordErrorMessage = "";
    }

	const login = async () => {
        resetAuthErrors();
        
        if(username.length === 0) {
            usernameIsError = true;
            usernameErrorMessage = "Username is required.";
            return;
        } else if(password.length === 0) {
            passwordIsError = true;
            passwordErrorMessage = "Password is required.";
            return;
        }

		authenticating = true;
		const response = await fetch("/auth/loginPost", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
			},
			credentials: "include",
			body: JSON.stringify({
				username: username,
				password: password,
			}),
		});

		if (response.status === 200) {
			// Go to dashboard... temp for now
			goto("/dashboard/1");
		} else if(response.status === 401) {
            usernameIsError = true;
            passwordIsError = true;
            passwordErrorMessage = "Invalid username and/or password.";
        } else if(response.status === 500) {
			passwordErrorMessage = "Something went wrong on our end. Please try again.";
		}
		authenticating = false;
	};
</script>

<svelte:head>
    <title>Log in - Pitstop</title>
</svelte:head>

<AuthenticationView>
    <span slot="heading" class="w-full text-center">
        <h2 class="text-lg md:text-xl font-medium p-1">Welcome back!</h2>
        <h3 class="text-xs text-slate-600">Please log in to your account to continue.</h3>
    </span>
    <span slot="body" class="w-full">
        <TextField label="Username" name="username" bind:value={username} isError={usernameIsError} errorMessage={usernameErrorMessage} />
        <PasswordField label="Password" name="password" bind:value={password} isError={passwordIsError} errorMessage={passwordErrorMessage} />
        
        <div class="flex flex-col w-full mt-6">
            <MainButton label={authenticating ? "Logging in..." : "Log in"} on:click={login} loading={authenticating} />
        </div>
        <div class="flex flex-col items-center w-full mt-4">
            <a href="/auth/register" class="text-amber-500 font-medium text-xs md:text-sm underline">Haven't got an account? Register</a>
        </div>
    </span>
</AuthenticationView>