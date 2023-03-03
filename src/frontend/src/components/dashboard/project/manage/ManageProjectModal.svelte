<script lang="ts">
    export let visible: boolean = true;
    export let category: number = 0;

	export let categories: Map<number, string> = new Map<number, string>([
		[0, "General"],
		[1, "Foobar"],
		[2, "Barfoo"],
		[3, "Baz"]
	]);
</script>

<div class="fixed w-full h-full bg-black/60 grid z-50 backdrop-blur-sm overflow-hidden" class:block={visible} class:hidden={!visible}>
    <div class="place-self-center justify-self-center w-[80%] h-[75%] bg-slate-800 border border-slate-700 text-slate-50 shadow-xl rounded-lg flex items-stretch overflow-hidden">
        <div class="w-72 bg-slate-800 p-4">
			<div class="border-b border-slate-600 py-4">
				<h1 class="font-medium text-xs">Manage</h1>
				<h2 class="text-xl font-serif">Foobar</h2>
			</div>
			<div class="py-4">
				<h1 class="font-medium text-xs">Category</h1>
			</div>
		</div>
		<div class="p-4 bg-slate-700/50 grow">
			bar
		</div>
    </div>
</div>