<script lang="ts">
    export let date: Date;
</script>

<div>
    <h1 class="text-lg my-2">On {date.toLocaleDateString("en-gb")}, you asked, and they answered:</h1>

    <table class="w-full border-collapse h-full rounded-md overflow-hidden my-4">
        <thead>
            <tr class="text-slate-50 text-left">
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold"></th>
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold">Question</th>
                <th class="sticky top-0 p-1 px-2 bg-slate-500 font-semibold">Feedback</th>
            </tr>
        </thead>
        <tbody>
            <tr class="bg-slate-100 border-t border-slate-200">
                <td class="font-medium text-lg p-2">
                    Satisfaction
                </td>
                <td class="text-md p-2 border-l border-r border-slate-200">
                    Do you feel satisfied with the current progression of the project?
                </td>
                <td class="font-medium p-2">
                    Somewhat Agree
                </td>
            </tr>
            <tr class="bg-slate-100 border-t border-slate-200">
                <td class="font-medium text-lg p-2">
                    Communication
                </td>
                <td class="text-md p-2 border-l border-r border-slate-200">
                    Have the developers maintained good, active communication with you throughout?
                </td>
                <td class="font-medium p-2">
                    Neither Agree nor Disagree
                </td>
            </tr>
            <tr class="bg-slate-100 border-t border-slate-200">
                <td class="font-medium text-lg p-2">
                    Don't Panic
                </td>
                <td class="text-md p-2 border-l border-r border-slate-200">
                    I don't know
                </td>
                <td class="font-medium p-2">
                    Strongly Disagree
                </td>
            </tr>
        </tbody>
    </table>

    <h1 class="text-lg my-2">Based on these survey results, here are some suggestions to try:</h1>
    <ul class="list-disc pl-6">
        <li class="px-2">Current project satisfaction is in a good state but could be improved further. Try to maintain a good dialogue with your customer, showing them progress and milestones, to keep them happy.</li>
        <li class="px-2">🥹.</li>
    </ul>
</div>