import{P as _}from"./index-9fe0b115.js";import{ao as S,ap as m,aq as f,ar as T,as as x,at as R,au as X,av as h,aw as y,ax as E,ay as Y,az as w,aA as C,_ as b}from"./index.js";import{d as B,f as n,a8 as o,_ as r,a9 as i,aa as s,a0 as l,l as p,E as k,ai as F,n as $,z as g}from"./vue-5c68ae35.js";import{Q as A}from"./antd-12f11a56.js";import"./useContentViewHeight-c2d18fda.js";import"./useWindowSizeFn-997fa1d0.js";import"./onMountedOrActivated-97a1bb6a.js";const P=["Fade","Scale","SlideY","ScrollY","SlideYReverse","ScrollYReverse","SlideX","ScrollX","SlideXReverse","ScrollXReverse","ScaleRotate","ExpandX","Expand"],N=P.map(e=>({label:e,value:e,key:e})),V=B({components:{Select:A,PageWrapper:_,FadeTransition:S,ScaleTransition:m,SlideYTransition:f,ScrollYTransition:T,SlideYReverseTransition:x,ScrollYReverseTransition:R,SlideXTransition:X,ScrollXTransition:h,SlideXReverseTransition:y,ScrollXReverseTransition:E,ScaleRotateTransition:Y,ExpandXTransition:w,ExpandTransition:C},setup(){const e=n("Fade"),a=n(!0);function t(){a.value=!1,setTimeout(()=>{a.value=!0},300)}return{options:N,value:e,start:t,show:a}}});const W={class:"flex"},z={class:"box"};function D(e,a,t,q,I,L){const c=o("Select"),u=o("a-button"),d=o("PageWrapper");return r(),i(d,{title:"动画组件示例"},{default:s(()=>[l("div",W,[p(c,{options:e.options,value:e.value,"onUpdate:value":a[0]||(a[0]=v=>e.value=v),placeholder:"选择动画",style:{width:"150px"}},null,8,["options","value"]),p(u,{type:"primary",class:"ml-4",onClick:e.start},{default:s(()=>[k(" start ")]),_:1},8,["onClick"])]),(r(),i(F(`${e.value}Transition`),null,{default:s(()=>[$(l("div",z,null,512),[[g,e.show]])]),_:1}))]),_:1})}const M=b(V,[["render",D],["__scopeId","data-v-147303b5"]]);export{M as default};