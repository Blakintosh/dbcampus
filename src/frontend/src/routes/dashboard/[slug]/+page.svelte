<script lang="ts">
	import Project from "./_project/Project.svelte";
	import { page } from '$app/stores';
</script>

<Project {...$page.data.projects[0]}/>