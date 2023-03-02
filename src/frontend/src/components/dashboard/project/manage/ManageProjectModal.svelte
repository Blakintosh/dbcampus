<script lang="ts">
    export let visible: boolean = true;
    export let category: number = 0;
</script>

<div class="fixed w-full h-full bg-black/60 grid z-50 backdrop-blur-sm overflow-hidden" class:block={visible} class:hidden={!visible}>
    <div class="place-self-center justify-self-center w-[80%] h-[75%] bg-slate-100 shadow-xl rounded-lg grid grid-cols-4 auto-rows-max overflow-hidden" style="grid-template-rows: auto 1fr;">
        <div class="col-span-1 p-4 text-lg font-medium bg-slate-600 text-slate-50 inline-block">
            Manage
        </div>
        <div class="col-span-3 p-4 text-lg bg-slate-600 text-slate-50">
            A software development project
        </div>
        <div class="col-span-1 bg-slate-700 p-4 h-full grow">
            Foooooo
        </div>
        <div class="col-span-3 shadow-inner shadow-slate-600 p-4 h-full grow">
            Barrrr
        </div>
    </div>
</div>