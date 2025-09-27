import React from "react";
import ReactDOM from "react-dom/client";
import { defineComponent, ref, onMounted, onUnmounted } from "vue";
import { Excalidraw } from "@excalidraw/excalidraw";
import "@excalidraw/excalidraw/index.css";

export default defineComponent({
    name: "WhiteBoard",
    setup() {
        const containerRef = ref<HTMLDivElement | null>(null);
        const rootRef = ref<ReactDOM.Root | null>(null);

        onMounted(() => {
            if (containerRef.value) {
                // 创建 React 应用程序
                const root = ReactDOM.createRoot(containerRef.value);

                // 渲染 React 组件
                root.render(
                    React.createElement(Excalidraw, {
                        theme: "light",
                        langCode: "zh-CN",  // 设置为中文界面
                        // 可以在这里添加其他 Excalidraw 属性
                    })
                );

                rootRef.value = root;
            }
        });

        onUnmounted(() => {
            if (rootRef.value) {
                rootRef.value.unmount();
            }
        });

        return () => (
            <div style={{ height: "100%", width: "100%" }}>
                <div
                    ref={containerRef}
                    style={{ height: "100%", width: "100%" }}
                />
            </div>
        );
    }
});
