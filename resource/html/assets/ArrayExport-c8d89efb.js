import{B as c}from"./BasicTable-9834a2fb.js";import"./componentMap-c5f962c4.js";import"./TableImg.vue_vue_type_style_index_0_lang-3cd25fdb.js";import"./index-c850e487.js";import{c as l,d,b as f,e as _,f as t,g as a}from"./data-ad9fd7c0.js";import{P as h}from"./index-9fe0b115.js";import{d as C,a8 as r,_ as F,a9 as x,aa as o,l as p,E as m}from"./vue-5c68ae35.js";import{_ as B}from"./index.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./antd-12f11a56.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./onMountedOrActivated-97a1bb6a.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./sortable.esm-4ae27e0b.js";import"./useContentViewHeight-c2d18fda.js";const S=C({components:{BasicTable:c,PageWrapper:h},setup(){function e(){f({data:t,header:a,filename:"二维数组方式导出excel.xlsx"})}function i(){_({sheetList:[{data:t,header:a,sheetName:"Sheet1"},{data:t,header:a,sheetName:"Sheet2"}],filename:"二维数组方式导出excel-多Sheet示例.xlsx"})}return{aoaToExcel:e,aoaToMultipleSheet:i,columns:l,data:d}}});function T(e,i,A,b,E,k){const u=r("a-button"),n=r("BasicTable"),s=r("PageWrapper");return F(),x(s,{title:"导出示例",content:"根据数组格式的数据进行导出"},{default:o(()=>[p(n,{title:"基础表格",columns:e.columns,dataSource:e.data},{toolbar:o(()=>[p(u,{onClick:e.aoaToExcel},{default:o(()=>[m(" 导出 ")]),_:1},8,["onClick"]),p(u,{onClick:e.aoaToMultipleSheet,danger:""},{default:o(()=>[m(" 导出多Sheet ")]),_:1},8,["onClick"])]),_:1},8,["columns","dataSource"])]),_:1})}const Z=B(S,[["render",T]]);export{Z as default};
