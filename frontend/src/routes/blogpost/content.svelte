<script lang="ts">
    import { CodeBlock } from '@skeletonlabs/skeleton';

    export let content:string;

    let parser:DOMParser = new DOMParser()
    let contentXML:Element = parser.parseFromString(content, "text/xml").getElementsByTagName("document")[0]
    let elements = contentXML.childNodes

    let test = "var x = 1;"
    console.log(elements[5].textContent)

    
    let parseCode = function(text:string|null):string {
        if(!text) return "hallo"
        
        let result:string = ""
        
        for(let i = 0; i < text.length; i++) {
            if(text[i] == ' ' && result[i-1] != ' ') {
                result += text[i]
                continue
            }
            
            if(text[i] == '\\') {
                if(text[i+1] == 'n') result += "\n"
                if(text[i+1] == 't') result += "\t"
                if(text[i+1] == 'n' || text[i+1] == 't') i++
                continue
            }
            
            result += text[i]
        }
        
        result = result.replace(/ +/g, ' ').replace(/(?:\n|^) /g, '\n');
        console.log(result.trim())
        return result.trim()
    }
    let mycode = parseCode("        int main() {\n        \t    Portal portal = Portal();\n        \t    portal.run(\"localhost\", 8080);\n}")
</script>

{#each elements as node}
    {#if node.nodeName == "header"}
        <h2 class="h2 text-2xl">{node.textContent}</h2>
    {:else if node.nodeName == "paragraph"}
        <p>{node.textContent}</p>
    {:else if node.nodeName == "codeblock"}
        <CodeBlock class="my-5" language="dart" lineNumbers={true} code={parseCode(node.textContent)}/>
    {:else if node.nodeName == "url"}
        <a class="chip variant-soft" href="//{node.textContent}">{node.textContent}</a>
    {/if}
{/each}