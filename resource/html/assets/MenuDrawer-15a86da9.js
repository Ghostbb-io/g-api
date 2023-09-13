var M=Object.defineProperty;var v=Object.getOwnPropertySymbols;var V=Object.prototype.hasOwnProperty,x=Object.prototype.propertyIsEnumerable;var I=(e,o,t)=>o in e?M(e,o,{enumerable:!0,configurable:!0,writable:!0,value:t}):e[o]=t,g=(e,o)=>{for(var t in o||(o={}))V.call(o,t)&&I(e,t,o[t]);if(v)for(var t of v(o))x.call(o,t)&&I(e,t,o[t]);return e};var b=(e,o,t)=>new Promise((a,s)=>{var d=l=>{try{i(t.next(l))}catch(p){s(p)}},u=l=>{try{i(t.throw(l))}catch(p){s(p)}},i=l=>l.done?a(l.value):Promise.resolve(l.value).then(d,u);i((t=t.apply(e,o)).next())});import{B as F}from"./BasicForm-33def721.js";import"./componentMap-c5f962c4.js";import{u as G}from"./useForm-61e3194d.js";import"./RadioButtonGroup-81fea792.js";import{I as k,_ as T}from"./index.js";import{m as h,d as C,f as P,u as w,c as $,a8 as B,_ as q,a9 as A,aa as N,l as O,ad as L}from"./vue-5c68ae35.js";import{K as S}from"./antd-12f11a56.js";import{B as j,a as K}from"./index-ae722a40.js";import{g as U,a as z,e as W}from"./menu-f81aa51d.js";const c=e=>e==="dir",r=e=>e==="menu",m=e=>e==="iframe",ie=[{title:"菜單名稱",dataIndex:"title",width:200,align:"left"},{title:"圖標",dataIndex:"icon",width:50,customRender:({record:e})=>h(k,{icon:e.icon})},{title:"組件",dataIndex:"component"},{title:"路由地址",dataIndex:"path"},{title:"路由名稱",dataIndex:"name"},{title:"排序",dataIndex:"sort",width:50},{title:"類型",dataIndex:"type",width:80,customRender:({record:e})=>{const o=e.type,t=o=="dir"?"blue":o=="menu"?"default":"orange",a=o=="dir"?"目錄":o=="menu"?"菜單":"Iframe";return h(S,{color:t},()=>a)}},{title:"狀態",dataIndex:"status",width:80,customRender:({record:e})=>{const o=e.status,t=o?"green":"red",a=o?"啟用":"停用";return h(S,{color:t},()=>a)}},{title:"創建時間",dataIndex:"createdAt",width:180},{title:"最後更新時間",dataIndex:"updatedAt",width:180}],se=[{field:"menuName",label:"菜單名稱",component:"Input",colProps:{span:8}},{field:"status",label:"狀態",component:"Select",componentProps:{options:[{label:"啟用",value:!0},{label:"停用",value:!1}]},colProps:{span:8}}],E=[{field:"type",label:"菜單類型",component:"RadioButtonGroup",defaultValue:"dir",componentProps:{options:[{label:"目錄",value:"dir"},{label:"菜單",value:"menu"},{label:"內嵌網址",value:"iframe"}]},colProps:{lg:24,md:24}},{field:"title",label:"title",component:"Input",required:!0},{field:"parentID",label:"上級目錄",component:"TreeSelect",defaultValue:0,componentProps:{fieldNames:{label:"title",key:"key",value:"key"},getPopupContainer:()=>document.body}},{field:"path",label:"路由地址",component:"Input",required:!0},{field:"name",label:"路由名稱",component:"Input",required:!0},{field:"redirect",label:"redirect",component:"Input",ifShow:({values:e})=>c(e.type)},{field:"sort",label:"排序",component:"InputNumber",required:!0},{field:"icon",label:"圖標",component:"IconPicker"},{field:"component",label:"組件路徑",component:"Input",ifShow:({values:e})=>r(e.type)},{field:"frameSrc",label:"網址",component:"Input",required:!0,ifShow:({values:e})=>m(e.type)},{field:"transitionName",label:"指定動畫",component:"Input",ifShow:({values:e})=>r(e.type)},{field:"currentActiveMenu",label:"激活菜單",component:"Input",ifShow:({values:e})=>r(e.type)},{field:"status",label:"狀態",component:"RadioButtonGroup",defaultValue:!0,componentProps:{options:[{label:"啟用",value:!0},{label:"停用",value:!1}]}},{field:"ignoreKeepAlive",label:"忽略緩存",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>r(e.type)||m(e.type)},{field:"affix",label:"固定標簽",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>r(e.type)||m(e.type)},{field:"hideBreadcrumb",label:"隱藏面包屑顯示",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]}},{field:"hideChildrenInMenu",label:"隱藏所有子菜單",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>c(e.type)},{field:"hideTab",label:"hideTab",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>r(e.type)||m(e.type)},{field:"hideMenu",label:"hideMenu",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>r(e.type)},{field:"ignoreRoute",label:"忽略路由",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>c(e.type)||r(e.type)},{field:"hidePathForChildren",label:"忽略本級path",component:"RadioButtonGroup",defaultValue:!1,componentProps:{options:[{label:"是",value:!0},{label:"否",value:!1}]},ifShow:({values:e})=>c(e.type)}],H=C({name:"MenuDrawer",components:{BasicDrawer:j,BasicForm:F},emits:["success","register"],setup(e,{emit:o}){const t=P(!0),a=P(0),[s,{resetFields:d,setFieldsValue:u,updateSchema:i,validate:l}]=G({labelWidth:100,schemas:E,showActionButtonGroup:!1,baseColProps:{lg:12,md:24}}),[p,{setDrawerProps:f,closeDrawer:R}]=K(n=>b(this,null,function*(){d(),f({confirmLoading:!1}),t.value=!!(n!=null&&n.isUpdate),a.value=n.id,w(t)&&u(g({},n.record));const y=yield U({dir:!0});y.unshift({key:0,title:"無",children:[]}),i({field:"parentID",componentProps:{treeData:y}})})),D=$(()=>w(t)?"編輯菜單":"新增菜單");function _(){return b(this,null,function*(){try{const n=yield l();f({confirmLoading:!0}),w(t)?yield W(a.value,n):yield z(n),R(),o("success")}finally{f({confirmLoading:!1})}})}return{registerDrawer:p,registerForm:s,getTitle:D,handleSubmit:_}}});function J(e,o,t,a,s,d){const u=B("BasicForm"),i=B("BasicDrawer");return q(),A(i,L(e.$attrs,{onRegister:e.registerDrawer,showFooter:"",title:e.getTitle,width:"50%",onOk:e.handleSubmit}),{default:N(()=>[O(u,{onRegister:e.registerForm},null,8,["onRegister"])]),_:1},16,["onRegister","title","onOk"])}const Q=T(H,[["render",J]]),ue=Object.freeze(Object.defineProperty({__proto__:null,default:Q},Symbol.toStringTag,{value:"Module"}));export{Q as M,ue as a,ie as c,se as s};