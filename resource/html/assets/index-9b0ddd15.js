import{I as k,_ as x}from"./index.js";import{B as y}from"./BasicForm-33def721.js";import"./componentMap-c5f962c4.js";import"./RadioButtonGroup-81fea792.js";import{searchList as F,actions as L,schemas as N}from"./data-a215bd35.js";import{P}from"./index-9fe0b115.js";import{d as T,a8 as a,_ as o,a9 as p,aa as r,a2 as s,l as _,a0 as n,$ as m,ab as d,a1 as l,F as u,ac as V,E as f}from"./vue-5c68ae35.js";import{K as W,a6 as c}from"./antd-12f11a56.js";import"./FormItem.vue_vue_type_script_lang-bb880c91.js";import"./helper-736f1539.js";import"./BasicForm.vue_vue_type_style_index_0_lang-4cc18aab.js";import"./index-cf0c6833.js";import"./useWindowSizeFn-997fa1d0.js";import"./useFormItem-16924efd.js";import"./uuid-31b8b5a4.js";import"./download-2647fd8d.js";import"./base64Conver-39fc0d26.js";import"./index-cf464bf9.js";import"./index-7f09b489.js";import"./useContentViewHeight-c2d18fda.js";import"./onMountedOrActivated-97a1bb6a.js";const b=T({components:{Icon:k,Tag:W,BasicForm:y,PageWrapper:P,[c.name]:c,[c.Item.name]:c.Item,AListItemMeta:c.Item.Meta},setup(){return{prefixCls:"list-search",list:F,actions:L,schemas:N}}});function w(e,A,D,E,M,z){const C=a("BasicForm"),$=a("Icon"),h=a("Tag"),v=a("a-list-item-meta"),B=a("a-list-item"),I=a("a-list"),g=a("PageWrapper");return o(),p(g,{class:s(e.prefixCls),title:"搜索列表"},{headerContent:r(()=>[_(C,{class:s(`${e.prefixCls}__header-form`),labelWidth:100,schemas:e.schemas,showActionButtonGroup:!1},null,8,["class","schemas"])]),default:r(()=>[n("div",{class:s(`${e.prefixCls}__container`)},[_(I,null,{default:r(()=>[(o(!0),m(u,null,d(e.list,i=>(o(),p(B,{key:i.id},{default:r(()=>[_(v,null,{description:r(()=>[n("div",{class:s(`${e.prefixCls}__content`)},l(i.content),3),n("div",{class:s(`${e.prefixCls}__action`)},[(o(!0),m(u,null,d(e.actions,t=>(o(),m("div",{key:t.icon,class:s(`${e.prefixCls}__action-item`)},[t.icon?(o(),p($,{key:0,class:s(`${e.prefixCls}__action-icon`),icon:t.icon,color:t.color},null,8,["class","icon","color"])):V("",!0),f(" "+l(t.text),1)],2))),128)),n("span",{class:s(`${e.prefixCls}__time`)},l(i.time),3)],2)]),title:r(()=>[n("p",{class:s(`${e.prefixCls}__title`)},l(i.title),3),n("div",null,[(o(!0),m(u,null,d(i.description,t=>(o(),p(h,{key:t,class:"mb-2"},{default:r(()=>[f(l(t),1)]),_:2},1024))),128))])]),_:2},1024)]),_:2},1024))),128))]),_:1})],2)]),_:1},8,["class"])}const ie=x(b,[["render",w],["__scopeId","data-v-708dfd24"]]);export{ie as default};