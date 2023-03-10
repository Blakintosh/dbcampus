<script lang="ts">
	import SmallButton from "../../../../common/SmallButton.svelte";

	export let enabled: boolean = false;
</script>

<div class="absolute bottom-8 left-[50%] translate-x-[-50%] w-[80%]">
	<div class="flex relative bg-slate-800 shadow-lg px-8 py-3 rounded-md border border-slate-700 items-center justify-between w-full {enabled ? "animate-fly-in" : "hidden"}">
		<div class="flex items-center gap-8 text-md font-medium">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" class="w-6 h-6 fill-amber-300/90"><!--! Font Awesome Pro 6.3.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. --><path d="M256 32c14.2 0 27.3 7.5 34.5 19.8l216 368c7.3 12.4 7.3 27.7 .2 40.1S486.3 480 472 480H40c-14.3 0-27.6-7.7-34.7-20.1s-7-27.8 .2-40.1l216-368C228.7 39.5 241.8 32 256 32zm0 128c-13.3 0-24 10.7-24 24V296c0 13.3 10.7 24 24 24s24-10.7 24-24V184c0-13.3-10.7-24-24-24zm32 224a32 32 0 1 0 -64 0 32 32 0 1 0 64 0z"/></svg>
			<span>You have unsaved changes.</span>
		</div>
		<div>
			<SmallButton type="secondary" label="Revert Changes"/>
			<SmallButton type="primary" label="Save Changes"/>
		</div>
	</div>
</div>