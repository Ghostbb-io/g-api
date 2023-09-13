import{B as T}from"./BasicTable-9834a2fb.js";import{T as S}from"./componentMap-c5f962c4.js";import"./TableImg.vue_vue_type_style_index_0_lang-3cd25fdb.js";import{u as B}from"./useTable-b33575de.js";import{g as w}from"./system-82075152.js";import{P as A}from"./index-9fe0b115.js";import y from"./DeptTree-17353b6e.js";import{b as D}from"./index-cf0c6833.js";import{A as F,c as k,s as E}from"./AccountModal-4070746a.js";import{aS as M,_ as I}from"./index.js";import{d as R,r as V,a8 as t,_,a9 as C,aa as n,l as a,E as P,ac as W}from"./vue-5c68ae35.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./antd-12f11a56.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./useWindowSizeFn-997fa1d0.js";import"./onMountedOrActivated-97a1bb6a.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./sortable.esm-4ae27e0b.js";import"./useContentViewHeight-c2d18fda.js";import"./index-daeb014c.js";import"./useContextMenu-839fbb72.js";const $=R({name:"AccountManagement",components:{BasicTable:T,PageWrapper:A,DeptTree:y,AccountModal:F,TableAction:S},setup(){const o=M(),[h,{openModal:r}]=D(),i=V({}),[b,{reload:c,updateTableDataRecord:l}]=B({title:"账号列表",api:w,rowKey:"id",columns:k,formConfig:{labelWidth:120,schemas:E,autoSubmitOnEnter:!0},useSearchForm:!0,showTableSetting:!0,bordered:!0,handleSearchInfoFn(e){return e},actionColumn:{width:120,title:"操作",dataIndex:"action"}});function s(){r(!0,{isUpdate:!1})}function u(e){r(!0,{record:e,isUpdate:!0})}function p(e){}function m({isUpdate:e,values:g}){e?l(g.id,g):c()}function d(e=""){i.deptId=e,c()}function f(e){o("/system/account_detail/"+e.id)}return{registerTable:b,registerModal:h,handleCreate:s,handleEdit:u,handleDelete:p,handleSuccess:m,handleSelect:d,handleView:f,searchInfo:i}}});function N(o,h,r,i,b,c){const l=t("DeptTree"),s=t("a-button"),u=t("TableAction"),p=t("BasicTable"),m=t("AccountModal"),d=t("PageWrapper");return _(),C(d,{dense:"",contentFullHeight:"",fixedHeight:"",contentClass:"flex"},{default:n(()=>[a(l,{class:"w-1/4 xl:w-1/5",onSelect:o.handleSelect},null,8,["onSelect"]),a(p,{onRegister:o.registerTable,class:"w-3/4 xl:w-4/5",searchInfo:o.searchInfo},{toolbar:n(()=>[a(s,{type:"primary",onClick:o.handleCreate},{default:n(()=>[P("新增账号")]),_:1},8,["onClick"])]),bodyCell:n(({column:f,record:e})=>[f.key==="action"?(_(),C(u,{key:0,actions:[{icon:"clarity:info-standard-line",tooltip:"查看用户详情",onClick:o.handleView.bind(null,e)},{icon:"clarity:note-edit-line",tooltip:"编辑用户资料",onClick:o.handleEdit.bind(null,e)},{icon:"ant-design:delete-outlined",color:"error",tooltip:"删除此账号",popConfirm:{title:"是否确认删除",placement:"left",confirm:o.handleDelete.bind(null,e)}}]},null,8,["actions"])):W("",!0)]),_:1},8,["onRegister","searchInfo"]),a(m,{onRegister:o.registerModal,onSuccess:o.handleSuccess},null,8,["onRegister","onSuccess"])]),_:1})}const fe=I($,[["render",N]]);export{fe as default};