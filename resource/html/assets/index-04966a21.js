import{u as s}from"./index-ae722a40.js";import _ from"./Drawer1-3e544129.js";import y from"./Drawer2-71bd1556.js";import E from"./Drawer3-33c06ead.js";import A from"./Drawer4-5fc29c1f.js";import R from"./Drawer5-966970f2.js";import{P as k}from"./index-9fe0b115.js";import{d as B,a8 as t,_ as $,a9 as P,aa as u,l as r,E as i}from"./vue-5c68ae35.js";import{Z as v}from"./antd-12f11a56.js";import{_ as W}from"./index.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./componentMap-c5f962c4.js";import"./useFormItem-16924efd.js";import"./RadioButtonGroup-81fea792.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./uuid-31b8b5a4.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./helper-736f1539.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./useForm-61e3194d.js";import"./useContentViewHeight-c2d18fda.js";import"./onMountedOrActivated-97a1bb6a.js";const b=B({components:{Alert:v,PageWrapper:k,Drawer1:_,Drawer2:y,Drawer3:E,Drawer4:A,Drawer5:R},setup(){const[e,{openDrawer:o,setDrawerProps:p}]=s(),[g,{openDrawer:f}]=s(),[d,{openDrawer:n}]=s(),[a,{openDrawer:m}]=s(),[w,{openDrawer:D}]=s();function l(){m(!0,{data:"content",info:"Info"})}function c(){o(),p({loading:!0}),setTimeout(()=>{p({loading:!1})},2e3)}return{register1:e,openDrawer1:o,register2:g,openDrawer2:f,register3:d,openDrawer3:n,register4:a,register5:w,openDrawer5:D,send:l,openDrawerLoading:c}}});function L(e,o,p,g,f,d){const n=t("Alert"),a=t("a-button"),m=t("Drawer1"),w=t("Drawer2"),D=t("Drawer3"),l=t("Drawer4"),c=t("Drawer5"),F=t("PageWrapper");return $(),P(F,{title:"抽屉组件使用示例"},{default:u(()=>[r(n,{message:"使用 useDrawer 进行抽屉操作","show-icon":""}),r(a,{type:"primary",class:"my-4",onClick:e.openDrawerLoading},{default:u(()=>[i(" 打开Drawer ")]),_:1},8,["onClick"]),r(n,{message:"内外同时控制显示隐藏","show-icon":""}),r(a,{type:"primary",class:"my-4",onClick:o[0]||(o[0]=C=>e.openDrawer2(!0))},{default:u(()=>[i(" 打开Drawer ")]),_:1}),r(n,{message:"自适应高度/显示footer","show-icon":""}),r(a,{type:"primary",class:"my-4",onClick:o[1]||(o[1]=C=>e.openDrawer3(!0))},{default:u(()=>[i(" 打开Drawer ")]),_:1}),r(n,{message:"内外数据交互","show-icon":""}),r(a,{type:"primary",class:"my-4",onClick:e.send},{default:u(()=>[i(" 打开Drawer并传递数据 ")]),_:1},8,["onClick"]),r(n,{message:"详情页模式","show-icon":""}),r(a,{type:"primary",class:"my-4",onClick:o[2]||(o[2]=C=>e.openDrawer5(!0))},{default:u(()=>[i(" 打开详情Drawer ")]),_:1}),r(m,{onRegister:e.register1},null,8,["onRegister"]),r(w,{onRegister:e.register2},null,8,["onRegister"]),r(D,{onRegister:e.register3},null,8,["onRegister"]),r(l,{onRegister:e.register4},null,8,["onRegister"]),r(c,{onRegister:e.register5},null,8,["onRegister"])]),_:1})}const ur=W(b,[["render",L]]);export{ur as default};
