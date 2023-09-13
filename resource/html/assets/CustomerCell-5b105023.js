import{B as $}from"./BasicTable-9834a2fb.js";import"./componentMap-c5f962c4.js";import{_ as B}from"./TableImg.vue_vue_type_style_index_0_lang-3cd25fdb.js";import{_}from"./index.js";import{a8 as r,_ as t,$ as l,a9 as i,aa as o,a0 as C,l as h,F as f,ab as k,ag as c,a2 as A,ac as b,d as z,E as g,a1 as u}from"./vue-5c68ae35.js";import{u as S}from"./useTable-b33575de.js";import{d as N}from"./table-9286518d.js";import{K as P,Y as V}from"./antd-12f11a56.js";import"./BasicForm-33def721.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./onMountedOrActivated-97a1bb6a.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./sortable.esm-4ae27e0b.js";const x={class:"img-div"};function D(e,I,w,T,v,L){const n=r("AImage"),m=r("PreviewGroup"),d=r("Badge");return e.imgList&&e.imgList.length?(t(),l("div",{key:0,class:A([e.prefixCls,"flex items-center mx-auto"]),style:c(e.getWrapStyle)},[e.simpleShow?(t(),i(d,{key:0,count:!e.showBadge||e.imgList.length==1?0:e.imgList.length},{default:o(()=>[C("div",x,[h(m,null,{default:o(()=>[(t(!0),l(f,null,k(e.imgList,(s,a)=>(t(),i(n,{key:s,width:e.size,style:c({display:a===0?"":"none !important"}),src:e.srcPrefix+s,fallback:e.fallback},null,8,["width","style","src","fallback"]))),128))]),_:1})])]),_:1},8,["count"])):(t(),i(m,{key:1},{default:o(()=>[(t(!0),l(f,null,k(e.imgList,(s,a)=>(t(),i(n,{key:s,width:e.size,style:c({marginLeft:a===0?0:e.margin+"px"}),src:e.srcPrefix+s,fallback:e.fallback},null,8,["width","style","src","fallback"]))),128))]),_:1}))],6)):b("",!0)}const E=_(B,[["render",D]]),F=[{title:"ID",dataIndex:"id"},{title:"头像",dataIndex:"avatar",width:100},{title:"分类",dataIndex:"category",width:80,align:"center",defaultHidden:!0},{title:"姓名",dataIndex:"name",width:120},{title:"图片列表1",dataIndex:"imgArr",helpMessage:["这是简单模式的图片列表","只会显示一张在表格中","但点击可预览多张图片"],width:140},{title:"照片列表2",dataIndex:"imgs",width:160},{title:"地址",dataIndex:"address"},{title:"编号",dataIndex:"no"},{title:"开始时间",dataIndex:"beginTime"},{title:"结束时间",dataIndex:"endTime"}],G=z({components:{BasicTable:$,TableImg:E,Tag:P,Avatar:V},setup(){const[e]=S({title:"自定义列内容",titleHelpMessage:"表格中所有头像、图片均为mock生成，仅用于演示图片占位",api:N,columns:F,bordered:!0,showTableSetting:!0});return{registerTable:e}}}),H={class:"p-4"};function M(e,I,w,T,v,L){const n=r("Tag"),m=r("Avatar"),d=r("TableImg"),s=r("BasicTable");return t(),l("div",H,[h(s,{onRegister:e.registerTable},{bodyCell:o(({column:a,record:p,text:y})=>[a.key==="id"?(t(),l(f,{key:0},[g(" ID: "+u(p.id),1)],64)):a.key==="no"?(t(),i(n,{key:1,color:"green"},{default:o(()=>[g(u(p.no),1)]),_:2},1024)):a.key==="avatar"?(t(),i(m,{key:2,size:60,src:p.avatar},null,8,["src"])):a.key==="imgArr"?(t(),i(d,{key:3,size:60,simpleShow:!0,imgList:y},null,8,["imgList"])):a.key==="imgs"?(t(),i(d,{key:4,size:60,imgList:y},null,8,["imgList"])):a.key==="category"?(t(),i(n,{key:5,color:"green"},{default:o(()=>[g(u(p.no),1)]),_:2},1024)):b("",!0)]),_:1},8,["onRegister"])])}const ce=_(G,[["render",M]]);export{ce as default};