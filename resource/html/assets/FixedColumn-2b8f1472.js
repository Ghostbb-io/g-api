import{B as m}from"./BasicTable-9834a2fb.js";import{T as c}from"./componentMap-c5f962c4.js";import"./TableImg.vue_vue_type_style_index_0_lang-3cd25fdb.js";import{u as p}from"./useTable-b33575de.js";import{d as u}from"./table-9286518d.js";import{d as f,a8 as a,_ as r,$ as b,l as _,aa as h,a9 as T,ac as w}from"./vue-5c68ae35.js";import{_ as x}from"./index.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./antd-12f11a56.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./onMountedOrActivated-97a1bb6a.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./sortable.esm-4ae27e0b.js";const C=[{title:"ID",dataIndex:"id",fixed:"left",width:280},{title:"姓名",dataIndex:"name",width:260},{title:"地址",dataIndex:"address"},{title:"编号",dataIndex:"no",width:300},{title:"开始时间",width:200,dataIndex:"beginTime"},{title:"结束时间",dataIndex:"endTime",width:200}],A=f({components:{BasicTable:m,TableAction:c},setup(){const[t]=p({title:"TableAction组件及固定列示例",api:u,columns:C,rowSelection:{type:"radio"},bordered:!0,actionColumn:{width:160,title:"Action",dataIndex:"action"}});function e(i){}function o(i){}return{registerTable:t,handleDelete:e,handleOpen:o}}}),I={class:"p-4"};function B(t,e,o,i,F,k){const l=a("TableAction"),d=a("BasicTable");return r(),b("div",I,[_(d,{onRegister:t.registerTable},{bodyCell:h(({column:s,record:n})=>[s.key==="action"?(r(),T(l,{key:0,actions:[{label:"删除",icon:"ic:outline-delete-outline",onClick:t.handleDelete.bind(null,n)}],dropDownActions:[{label:"启用",popConfirm:{title:"是否启用？",confirm:t.handleOpen.bind(null,n)}}]},null,8,["actions","dropDownActions"])):w("",!0)]),_:1},8,["onRegister"])])}const X=x(A,[["render",B]]);export{X as default};