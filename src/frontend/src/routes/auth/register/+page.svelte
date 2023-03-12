<script lang="ts">
	import MainButton from "../../../components/common/ActionButton.svelte";
    import AuthenticationView from "../../../components/auth/AuthenticationView.svelte";
	import TextField from "../../../components/auth/TextField.svelte";
	import PasswordField from "../../../components/auth/PasswordField.svelte";
	import { goto } from "$app/navigation";

	let username: string = "";
	let password: string = "";
    let confirmPassword: string = "";

	let usernameIsError = false;
	let usernameErrorMessage = "";

	let passwordIsError = false;
	let passwordErrorMessage = "";

	let confirmPasswordIsError = false;
	let confirmPasswordErrorMessage = "";

	let authenticating: boolean = false;

	const resetAuthErrors = () => {
        usernameIsError = false;
        usernameErrorMessage = "";
        passwordIsError = false;
        passwordErrorMessage = "";
		confirmPasswordIsError = false;
		confirmPasswordErrorMessage = "";
    }

	const register = async () => {
        resetAuthErrors();
        
        if(username.length === 0) {
            usernameIsError = true;
            usernameErrorMessage = "Username is required.";
            return;
        } else if(password.length === 0) {
            passwordIsError = true;
            passwordErrorMessage = "Password is required.";
            return;
        } else if(password !== confirmPassword) {
            passwordIsError = true;
			confirmPasswordIsError = true;
			confirmPasswordErrorMessage = "Passwords do not match.";
            return;
        }

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

		authenticating = false;

        if(response.status === 200) {
            goto("/auth/login");
        }
	};
</script>


<svelte:head>
    <title>Register - Pitstop</title>
</svelte:head>

<AuthenticationView>
    <span slot="heading" class="w-full text-center">
        <h2 class="text-lg md:text-xl font-medium p-1">Registration</h2>
    </span>
    <span slot="body" class="w-full">
        <TextField label="Username" name="username" bind:value={username} isError={usernameIsError} errorMessage={usernameErrorMessage} />
        <PasswordField label="Password" name="password" bind:value={password} isError={passwordIsError} errorMessage={passwordErrorMessage} />
        <PasswordField label="Confirm Password" name="confirmPassword" bind:value={confirmPassword} isError={confirmPasswordIsError} errorMessage={confirmPasswordErrorMessage} />
        
        <div class="flex flex-col w-full mt-6">
            <MainButton label={authenticating ? "Registering..." : "Register"} on:click={register} loading={authenticating} />
        </div>
        <div class="flex flex-col items-center w-full mt-4">
            <a href="/auth/login" class="text-amber-500 font-medium text-xs md:text-sm underline">or go to Log in</a>
        </div>
    </span>
</AuthenticationView>