import{B as m}from"./BasicForm-33def721.js";import"./componentMap-c5f962c4.js";import"./RadioButtonGroup-81fea792.js";import{C as u,b as l,_ as c}from"./index.js";import{T as d}from"./index-e1293ec7.js";import{P as C}from"./index-9fe0b115.js";import{d as f,m as _,a8 as o,_ as h,a9 as b,aa as r,l as s}from"./vue-5c68ae35.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./antd-12f11a56.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./onMountedOrActivated-97a1bb6a.js";import"./useContentViewHeight-c2d18fda.js";const B=[{field:"title",component:"Input",label:"title",defaultValue:"defaultValue",rules:[{required:!0}]},{field:"tinymce",component:"Input",label:"tinymce",defaultValue:"defaultValue",rules:[{required:!0}],render:({model:e,field:t})=>_(d,{value:e[t],onChange:a=>{e[t]=a}})}],g=f({components:{BasicForm:m,CollapseContainer:u,PageWrapper:C},setup(){const{createMessage:e}=l();return{schemas:B,handleSubmit:t=>{e.success("click search,values:"+JSON.stringify(t))}}}});function P(e,t,a,S,V,y){const n=o("BasicForm"),i=o("CollapseContainer"),p=o("PageWrapper");return h(),b(p,{title:"富文本嵌入表单示例"},{default:r(()=>[s(i,{title:"富文本表单"},{default:r(()=>[s(n,{labelWidth:100,schemas:e.schemas,actionColOptions:{span:24},baseColProps:{span:24},onSubmit:e.handleSubmit},null,8,["schemas","onSubmit"])]),_:1})]),_:1})}const K=c(g,[["render",P]]);export{K as default};
