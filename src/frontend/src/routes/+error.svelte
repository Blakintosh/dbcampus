<script>
	import { page } from '$app/stores';

	const emojis = {
		// TODO add the rest!
		420: '🫠',
		500: '💥'
	};
</script>


<main class="w-[100vw] h-[100vh] grid bg-slate-100 shadow-inner">
    <!-- <h1>{$page.status} {$page.error.message}</h1>
    <span style="font-size: 10em">
        {emojis[$page.status] ?? emojis[500]}
    </span> -->

    <div class="place-self-center justify-self-center flex flex-col items-center gap-16">
        <div class="text-center">
            <h1 class="text-9xl font-bold font-serif text-amber-500 mb-4 drop-shadow-md">{$page.status}</h1>
            <h3 class="text-4xl mb-2">{$page.error?.message}</h3>
            <h4 class="text-xl">Sorry about that.</h4>
        </div>
        <a href="/">
            <img src="/pitstop.png" alt="Pitstop Logo" class="w-[25vw]">
        </a>
    </div>
</main>