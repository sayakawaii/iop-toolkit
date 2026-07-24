import { title } from "@/components/primitives";
import DefaultLayout from "@/layouts/default";
import {Form,Input,Button} from "@heroui/react";
import { button as buttonStyles } from "@heroui/theme";
import React, { useEffect} from "react";
import axios from "axios";
import { API_BASE } from "@/config/api";

export default function SequenceTracerPage() {
    const [submitted, setSubmitted] = React.useState<null | { [k: string]: FormDataEntryValue }>(null);
    const mermaidRef = React.useRef<HTMLDivElement | null>(null);
    const [mermaidCode, setMermaidCode] = React.useState<string | null>(null);
    const [isEmptyResult, setIsEmptyResult] = React.useState<boolean | null>(null);
    const [loading, setLoading] = React.useState(false);
    const [error, setError] = React.useState<string | null>(null);

    const onSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const data = Object.fromEntries(new FormData(e.currentTarget));
        setSubmitted(data);
    };

    const baseUrl = API_BASE + "/api/sequencetracer";
    // upload request and kick off polling for returned requestKey(s)
    useEffect(() => {
        if (!submitted) return;

        let cancelled = false;
        const fetchData = async () => {
            setLoading(true);
            setError(null);
            try {
                const formData = new FormData();
                for (const key in submitted) {
                    const value = submitted[key];
                    if (value instanceof File) {
                        formData.append(key, value);
                    } else {
                        formData.append(key, value.toString());
                    }
                }

                // POST files to backend
                const response = await axios.post(`${baseUrl}/request`, formData, {
                    headers: {
                        'Content-Type': 'multipart/form-data',
                    },
                });

                if (cancelled) return;
                console.log('Response data:', response.data);
                // 处理后端返回：{ mermaidData: string, isEmpty: bool }
                const { mermaidData, isEmpty } = response.data ?? {};
                setMermaidCode(mermaidData ?? null);
                setIsEmptyResult(typeof isEmpty === "boolean" ? isEmpty : null);
            } catch (err: any) {
                console.error('Error submitting form:', err);
                setError(err?.message ?? String(err));
                setMermaidCode(null);
                setIsEmptyResult(null);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
        return () => {
            cancelled = true;
        };
    }, [submitted]);
 
    // 当 mermaidCode 更新时，在页面内渲染 mermaid 图（采用 mermaid 的 DOM init 流程，避免 SSR 导入问题）
    useEffect(() => {
        if (!mermaidCode) return;
        let cancelled = false;
        (async () => {
            const container = mermaidRef.current;
            if (!container) return;
            try {
                const m = await import("mermaid");
                // mermaid 的 default export 在某些打包器下存在于 default
                const mermaid = (m as any).default ?? m;
                mermaid.initialize({
                    startOnLoad: false,
                    securityLevel: "loose", // 根据需要调整
                });
                // 清空旧内容，创建一个 .mermaid 节点并将原始 mermaid 文本作为 textContent 插入
                // (不要把 mermaid 文本包裹成 HTML，否则 mermaid 会把 <div> 一起当成文本导致 UnknownDiagramError)
                container.innerHTML = "";
                const node = document.createElement("div");
                node.className = "mermaid";
                node.textContent = mermaidCode;
                container.appendChild(node);
                // 让 mermaid 扫描并渲染刚刚插入的节点
                mermaid.init(undefined, node);
            } catch (e) {
                if (!cancelled) {
                    console.error("Failed to render mermaid:", e);
                    setError("Failed to render diagram");
                }
            }
        })();
        return () => {
            cancelled = true;
        };
    }, [mermaidCode]);
   return (
     <DefaultLayout fullWidth>
       <section className="flex flex-col items-center justify-center gap-4 py-8 md:py-10">
         <div className="inline-block max-w-lg text-center justify-center">
           <span className={title({size:"sm"})}>Sequence Tracer&nbsp;</span>
         </div>
         <Form className="w-full max-w-xs" onSubmit={onSubmit}>
             <div className="w-full mt-8 justify-center">
                 <Input name="sequencetracerFile" className="justify-center" color="primary" type="file" multiple />
             </div>
             <div className="flex gap-10 justify-center">
                 <Button type="submit" className="bg-linear-to-tr from-pink-500 to-yellow-500 text-white shadow-lg" radius="full">
                     Start
                 </Button>
                 <Button type="reset" className={buttonStyles({
                     color: "default",
                     radius: "full",
                     variant: "shadow",
                     })}>
                     Reset
                 </Button>
             </div>
         </Form>
        <div className="w-full mt-6">
            {loading && <div>Rendering...</div>}
            {error && <div className="text-red-600">Error: {error}</div>}
            {isEmptyResult && <div className="text-gray-600">No sequence data produced (empty result).</div>}
            {/* mermaid output container */}
            <div ref={mermaidRef} className="mx-auto mt-4" />
        </div>
       </section>
     </DefaultLayout>
   );
 }
