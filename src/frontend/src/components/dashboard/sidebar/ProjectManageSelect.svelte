<script lang="ts">
	import { ProjectManageCategory } from "../../../util/models";
	import ProjectTabButton from "./ProjectTabButton.svelte";
    import { modalCategory, modalVisible } from "../../../util/stores";

    const openModal = (category: ProjectManageCategory) => {
        modalVisible.set(true);
        modalCategory.set(category);
    }
</script>

<ul class="mb-6">
	<ProjectTabButton label="Update Data" on:click={() => openModal(ProjectManageCategory.General)}>
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M471.6 21.7c-21.9-21.9-57.3-21.9-79.2 0L362.3 51.7l97.9 97.9 30.1-30.1c21.9-21.9 21.9-57.3 0-79.2L471.6 21.7zm-299.2 220c-6.1 6.1-10.8 13.6-13.5 21.9l-29.6 88.8c-2.9 8.6-.6 18.1 5.8 24.6s15.9 8.7 24.6 5.8l88.8-29.6c8.2-2.8 15.7-7.4 21.9-13.5L437.7 172.3 339.7 74.3 172.4 241.7zM96 64C43 64 0 107 0 160V416c0 53 43 96 96 96H352c53 0 96-43 96-96V320c0-17.7-14.3-32-32-32s-32 14.3-32 32v96c0 17.7-14.3 32-32 32H96c-17.7 0-32-14.3-32-32V160c0-17.7 14.3-32 32-32h96c17.7 0 32-14.3 32-32s-14.3-32-32-32H96z"/></svg>
	</ProjectTabButton>
	<ProjectTabButton label="Issue a Survey" on:click={() => openModal(ProjectManageCategory.Surveys)}>
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M192 0c-41.8 0-77.4 26.7-90.5 64H64C28.7 64 0 92.7 0 128V448c0 35.3 28.7 64 64 64H320c35.3 0 64-28.7 64-64V128c0-35.3-28.7-64-64-64H282.5C269.4 26.7 233.8 0 192 0zm0 64a32 32 0 1 1 0 64 32 32 0 1 1 0-64zM105.8 229.3c7.9-22.3 29.1-37.3 52.8-37.3h58.3c34.9 0 63.1 28.3 63.1 63.1c0 22.6-12.1 43.5-31.7 54.8L216 328.4c-.2 13-10.9 23.6-24 23.6c-13.3 0-24-10.7-24-24V314.5c0-8.6 4.6-16.5 12.1-20.8l44.3-25.4c4.7-2.7 7.6-7.7 7.6-13.1c0-8.4-6.8-15.1-15.1-15.1H158.6c-3.4 0-6.4 2.1-7.5 5.3l-.4 1.2c-4.4 12.5-18.2 19-30.6 14.6s-19-18.2-14.6-30.6l.4-1.2zM160 416a32 32 0 1 1 64 0 32 32 0 1 1 -64 0z"/></svg>
	</ProjectTabButton>
	<ProjectTabButton label="Project Settings" on:click={() => openModal(ProjectManageCategory.Settings)}>
		<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M481.9 166.6c3.2 8.7 .5 18.4-6.4 24.6l-30.9 28.1c-7.7 7.1-11.4 17.5-10.9 27.9c.1 2.9 .2 5.8 .2 8.8s-.1 5.9-.2 8.8c-.5 10.5 3.1 20.9 10.9 27.9l30.9 28.1c6.9 6.2 9.6 15.9 6.4 24.6c-4.4 11.9-9.7 23.3-15.8 34.3l-4.7 8.1c-6.6 11-14 21.4-22.1 31.2c-5.9 7.2-15.7 9.6-24.5 6.8l-39.7-12.6c-10-3.2-20.8-1.1-29.7 4.6c-4.9 3.1-9.9 6.1-15.1 8.7c-9.3 4.8-16.5 13.2-18.8 23.4l-8.9 40.7c-2 9.1-9 16.3-18.2 17.8c-13.8 2.3-28 3.5-42.5 3.5s-28.7-1.2-42.5-3.5c-9.2-1.5-16.2-8.7-18.2-17.8l-8.9-40.7c-2.2-10.2-9.5-18.6-18.8-23.4c-5.2-2.7-10.2-5.6-15.1-8.7c-8.8-5.7-19.7-7.8-29.7-4.6L69.1 425.9c-8.8 2.8-18.6 .3-24.5-6.8c-8.1-9.8-15.5-20.2-22.1-31.2l-4.7-8.1c-6.1-11-11.4-22.4-15.8-34.3c-3.2-8.7-.5-18.4 6.4-24.6l30.9-28.1c7.7-7.1 11.4-17.5 10.9-27.9c-.1-2.9-.2-5.8-.2-8.8s.1-5.9 .2-8.8c.5-10.5-3.1-20.9-10.9-27.9L8.4 191.2c-6.9-6.2-9.6-15.9-6.4-24.6c4.4-11.9 9.7-23.3 15.8-34.3l4.7-8.1c6.6-11 14-21.4 22.1-31.2c5.9-7.2 15.7-9.6 24.5-6.8l39.7 12.6c10 3.2 20.8 1.1 29.7-4.6c4.9-3.1 9.9-6.1 15.1-8.7c9.3-4.8 16.5-13.2 18.8-23.4l8.9-40.7c2-9.1 9-16.3 18.2-17.8C213.3 1.2 227.5 0 242 0s28.7 1.2 42.5 3.5c9.2 1.5 16.2 8.7 18.2 17.8l8.9 40.7c2.2 10.2 9.4 18.6 18.8 23.4c5.2 2.7 10.2 5.6 15.1 8.7c8.8 5.7 19.7 7.7 29.7 4.6l39.7-12.6c8.8-2.8 18.6-.3 24.5 6.8c8.1 9.8 15.5 20.2 22.1 31.2l4.7 8.1c6.1 11 11.4 22.4 15.8 34.3zM242 336a80 80 0 1 0 0-160 80 80 0 1 0 0 160z"/></svg>
	</ProjectTabButton>
</ul>