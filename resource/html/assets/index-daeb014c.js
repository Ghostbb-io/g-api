var Oe=Object.defineProperty,Pe=Object.defineProperties;var $e=Object.getOwnPropertyDescriptors;var se=Object.getOwnPropertySymbols;var me=Object.prototype.hasOwnProperty,Ue=Object.prototype.propertyIsEnumerable;var re=(e,a,l)=>a in e?Oe(e,a,{enumerable:!0,configurable:!0,writable:!0,value:l}):e[a]=l,O=(e,a)=>{for(var l in a||(a={}))me.call(a,l)&&re(e,l,a[l]);if(se)for(var l of se(a))Ue.call(a,l)&&re(e,l,a[l]);return e},ie=(e,a)=>Pe(e,$e(a));var ue=(e,a,l)=>new Promise((k,_)=>{var A=K=>{try{p(l.next(K))}catch(E){_(E)}},s=K=>{try{p(l.throw(K))}catch(E){_(E)}},p=K=>K.done?k(K.value):Promise.resolve(K.value).then(A,s);p((l=l.apply(e,a)).next())});import{ao as je,d as Ce,f as Ae,ah as He,c as M,w as $,_ as I,$ as U,u,N as Ve,ac as P,a9 as q,aa as j,E as oe,a1 as de,a2 as he,l as L,F as Q,ab as Xe,aj as Ye,ak as Re,A as ze,m as Ge,r as fe,H as R,h as H,o as Je,n as ye,z as pe,ad as qe,i as We}from"./vue-5c68ae35.js";import{aD as Qe,k as Ze,u as et,a6 as tt,I as Z,aE as ke,L as nt,m as W,aF as at,ae as lt,aG as ct,a3 as st,a4 as rt,S as it,O as ge,x as ut}from"./index.js";import{aE as ot,aF as dt,aG as ht,l as ft,aH as yt,aI as pt,D as kt,aJ as gt,d as ee,aK as St,o as Kt,A as xt,aa as Ct,E as Se,a4 as At}from"./antd-12f11a56.js";import{u as Lt}from"./useContextMenu-839fbb72.js";function te(e,a){return a?typeof a=="string"?` ${e}--${a}`:Array.isArray(a)?a.reduce((l,k)=>l+te(e,k),""):Object.keys(a).reduce((l,k)=>l+(a[k]?te(e,k):""),""):""}function vt(e){return(a,l)=>(a&&typeof a!="string"&&(l=a,a=""),a=a?`${e}__${a}`:e,`${a}${te(a,l)}`)}function Le(e){return[vt(`${Qe}-${e}`)]}const Ke=Symbol(),xe=Symbol();function Et(e,a){if(!dt(e)||e[xe])return e;const{values:l,required:k,default:_,type:A,validator:s}=e,p=l||s?K=>{let E=!1,v=[];if(l&&(v=[...l,_],E||(E=v.includes(K))),s&&(E||(E=s(K))),!E&&v.length>0){const h=[...new Set(v)].map(t=>JSON.stringify(t)).join(", ");je(`Invalid prop: validation failed${a?` for prop "${a}"`:""}. Expected one of [${h}], got value ${JSON.stringify(K)}.`)}return E}:void 0;return{type:typeof A=="object"&&Object.getOwnPropertySymbols(A).includes(Ke)?A[Ke]:A,required:!!k,default:_,validator:p,[xe]:!0}}const bt=e=>ot(Object.entries(e).map(([a,l])=>[a,Et(l,a)]));var T=(e=>(e[e.SELECT_ALL=0]="SELECT_ALL",e[e.UN_SELECT_ALL=1]="UN_SELECT_ALL",e[e.EXPAND_ALL=2]="EXPAND_ALL",e[e.UN_EXPAND_ALL=3]="UN_EXPAND_ALL",e[e.CHECK_STRICTLY=4]="CHECK_STRICTLY",e[e.CHECK_UN_STRICTLY=5]="CHECK_UN_STRICTLY",e))(T||{});const Tt=["update:expandedKeys","update:selectedKeys","update:value","change","check","update:searchValue"],_t=bt({value:{type:[Object,Array]},renderIcon:{type:Function},helpMessage:{type:[String,Array],default:""},title:{type:String,default:""},toolbar:Boolean,search:Boolean,searchValue:{type:String,default:""},checkStrictly:Boolean,clickRowToExpand:{type:Boolean,default:!1},checkable:Boolean,defaultExpandLevel:{type:[String,Number],default:""},defaultExpandAll:Boolean,fieldNames:{type:Object},treeData:{type:Array},actionList:{type:Array,default:()=>[]},expandedKeys:{type:Array,default:()=>[]},selectedKeys:{type:Array,default:()=>[]},checkedKeys:{type:[Array,Object],default:()=>[]},beforeRightClick:{type:Function,default:void 0},rightMenuList:{type:Array},filterFn:{type:Function,default:void 0},highlight:{type:[Boolean,String],default:!1},expandOnSearch:Boolean,checkOnSearch:Boolean,selectedOnSearch:Boolean,loading:{type:Boolean,default:!1},treeWrapperClassName:String}),Bt={key:2,class:"flex items-center flex-1 cursor-pointer justify-self-stretch"},Ft=Ce({__name:"TreeHeader",props:{helpMessage:{type:[String,Array],default:""},title:{type:String,default:""},toolbar:{type:Boolean,default:!1},checkable:{type:Boolean,default:!1},search:{type:Boolean,default:!1},searchText:{type:String,default:""},checkAll:{type:Function,default:void 0},expandAll:{type:Function,default:void 0}},emits:["strictly-change","search"],setup(e,{emit:a}){const l=e,k=Ae(""),[_]=Le("tree-header"),A=He(),{t:s}=Ze(),p=M(()=>["mr-1","w-full",{"ml-5":A.headerTitle||l.title}]),K=M(()=>{const{checkable:t}=l,c=[{label:s("component.tree.expandAll"),value:T.EXPAND_ALL},{label:s("component.tree.unExpandAll"),value:T.UN_EXPAND_ALL,divider:t}];return t?[{label:s("component.tree.selectAll"),value:T.SELECT_ALL},{label:s("component.tree.unSelectAll"),value:T.UN_SELECT_ALL,divider:t},...c,{label:s("component.tree.checkStrictly"),value:T.CHECK_STRICTLY},{label:s("component.tree.checkUnStrictly"),value:T.CHECK_UN_STRICTLY}]:c});function E(t){var i,d,o,r;const{key:c}=t;switch(c){case T.SELECT_ALL:(i=l.checkAll)==null||i.call(l,!0);break;case T.UN_SELECT_ALL:(d=l.checkAll)==null||d.call(l,!1);break;case T.EXPAND_ALL:(o=l.expandAll)==null||o.call(l,!0);break;case T.UN_EXPAND_ALL:(r=l.expandAll)==null||r.call(l,!1);break;case T.CHECK_STRICTLY:a("strictly-change",!1);break;case T.CHECK_UN_STRICTLY:a("strictly-change",!0);break}}function v(t){a("search",t)}const h=et(v,200);return $(()=>k.value,t=>{h(t)}),$(()=>l.searchText,t=>{t!==k.value&&(k.value=t)}),(t,c)=>(I(),U("div",{class:he([u(_)(),"flex px-2 py-1.5 items-center"])},[u(A).headerTitle?Ve(t.$slots,"headerTitle",{key:0}):P("",!0),!u(A).headerTitle&&t.title?(I(),q(u(tt),{key:1,helpMessage:t.helpMessage},{default:j(()=>[oe(de(t.title),1)]),_:1},8,["helpMessage"])):P("",!0),t.search||t.toolbar?(I(),U("div",Bt,[t.search?(I(),U("div",{key:0,class:he(p.value)},[L(u(ht),{placeholder:u(s)("common.searchText"),size:"small",allowClear:"",value:k.value,"onUpdate:value":c[0]||(c[0]=i=>k.value=i)},null,8,["placeholder","value"])],2)):P("",!0),t.toolbar?(I(),q(u(kt),{key:1,onClick:c[1]||(c[1]=ze(()=>{},["prevent"]))},{overlay:j(()=>[L(u(ft),{onClick:E},{default:j(()=>[(I(!0),U(Q,null,Xe(K.value,i=>(I(),U(Q,{key:i.value},[L(u(yt),Ye(Re({key:i.value})),{default:j(()=>[oe(de(i.label),1)]),_:2},1040),i.divider?(I(),q(u(pt),{key:0})):P("",!0)],64))),128))]),_:1})]),default:j(()=>[L(Z,{icon:"ion:ellipsis-vertical"})]),_:1})):P("",!0)])):P("",!0)],2))}}),wt=({icon:e})=>e?gt(e)?Ge(Z,{icon:e,class:"mr-1"}):Z:null;function Dt(e,a){function l(h){const t=[],c=h||u(e),{key:i,children:d}=u(a);if(!d||!i)return t;for(let o=0;o<c.length;o++){const r=c[o];t.push(r[i]);const y=r[d];y&&y.length&&t.push(...l(y))}return t}function k(h){const t=[],c=h||u(e),{key:i,children:d}=u(a);if(!d||!i)return t;for(let o=0;o<c.length;o++){const r=c[o];r.disabled!==!0&&r.selectable!==!1&&t.push(r[i]);const y=r[d];y&&y.length&&t.push(...k(y))}return t}function _(h,t){const c=[],i=t||u(e),{key:d,children:o}=u(a);if(!o||!d)return c;for(let r=0;r<i.length;r++){const y=i[r],b=y[o];h===y[d]?(c.push(y[d]),b&&b.length&&c.push(...l(b))):b&&b.length&&c.push(..._(h,b))}return c}function A(h,t,c){if(!h)return;const i=c||u(e),{key:d,children:o}=u(a);if(!(!o||!d))for(let r=0;r<i.length;r++){const y=i[r],b=y[o];if(y[d]===h){i[r]=O(O({},i[r]),t);break}else b&&b.length&&A(h,t,y[o])}}function s(h=1,t,c=1){if(!h)return[];const i=[],d=t||u(e)||[];for(let o=0;o<d.length;o++){const r=d[o],{key:y,children:b}=u(a),z=y?r[y]:"",m=b?r[b]:[];i.push(z),m&&m.length&&c<h&&(c+=1,i.push(...s(h,m,c)))}return i}function p({parentKey:h=null,node:t,push:c="push"}){const i=ee(u(e));if(!h){i[c](t),e.value=i;return}const{key:d,children:o}=u(a);!o||!d||(ke(i,r=>{if(r[d]===h)return r[o]=r[o]||[],r[o][c](t),!0}),e.value=i)}function K({parentKey:h=null,list:t,push:c="push"}){const i=ee(u(e));if(!(!t||t.length<1))if(h){const{key:d,children:o}=u(a);if(!o||!d)return;ke(i,r=>{if(r[d]===h){r[o]=r[o]||[];for(let y=0;y<t.length;y++)r[o][c](t[y]);return e.value=i,!0}})}else{for(let d=0;d<t.length;d++)i[c](t[d]);e.value=i;return}}function E(h,t){if(!h)return;const c=t||u(e),{key:i,children:d}=u(a);if(!(!d||!i))for(let o=0;o<c.length;o++){const r=c[o],y=r[d];if(r[i]===h){c.splice(o,1);break}else y&&y.length&&E(h,r[d])}}function v(h,t,c){if(!h&&h!==0)return null;const i=t||u(e),{key:d,children:o}=u(a);if(d)return i.forEach(r=>{if(c!=null&&c.key||(c==null?void 0:c.key)===0)return c;if(r[d]===h){c=r;return}r[o]&&r[o].length&&(c=v(h,r[o],c))}),c||null}return{deleteNodeByKey:E,insertNodeByKey:p,insertNodesByKey:K,filterByLevel:s,updateNodeByKey:A,getAllKeys:l,getChildrenKeys:_,getEnabledKeys:k,getSelectedNode:v}}function Nt(e){return typeof e=="function"||Object.prototype.toString.call(e)==="[object Object]"&&!We(e)}const mt=Ce({name:"BasicTree",inheritAttrs:!1,props:_t,emits:Tt,setup(e,{attrs:a,slots:l,emit:k,expose:_}){const[A]=Le("tree"),s=fe({checkStrictly:e.checkStrictly,expandedKeys:e.expandedKeys||[],selectedKeys:e.selectedKeys||[],checkedKeys:e.checkedKeys||[]}),p=fe({startSearch:!1,searchText:"",searchData:[]}),K=Ae([]),[E]=Lt(),v=M(()=>{const{fieldNames:n}=e;return O({children:"children",title:"title",key:"key"},n)}),h=M(()=>{let n=ie(O(O({blockNode:!0},a),e),{expandedKeys:s.expandedKeys,selectedKeys:s.selectedKeys,checkedKeys:s.checkedKeys,checkStrictly:s.checkStrictly,fieldNames:u(v),"onUpdate:expandedKeys":f=>{s.expandedKeys=f,k("update:expandedKeys",f)},"onUpdate:selectedKeys":f=>{s.selectedKeys=f,k("update:selectedKeys",f)},onCheck:(f,x)=>{let g=R(s.checkedKeys);if(nt(g)&&p.startSearch){const C=x.node.eventKey;g=St(g,z(C)),x.checked&&g.push(C),s.checkedKeys=g}else s.checkedKeys=f;const S=R(s.checkedKeys);k("update:value",S),k("check",S,x)},onRightClick:be});return Kt(n,"treeData","class")}),t=M(()=>p.startSearch?p.searchData:u(K)),c=M(()=>!t.value||t.value.length===0),{deleteNodeByKey:i,insertNodeByKey:d,insertNodesByKey:o,filterByLevel:r,updateNodeByKey:y,getAllKeys:b,getChildrenKeys:z,getEnabledKeys:m,getSelectedNode:ve}=Dt(K,v);function Ee(n,f){return!f&&e.renderIcon&&W(e.renderIcon)?e.renderIcon(n):f}function be(x){return ue(this,arguments,function*({event:n,node:f}){var w;const{rightMenuList:g=[],beforeRightClick:S}=e;let C={event:n,items:[]};if(S&&W(S)){let B=yield S(f,n);Array.isArray(B)?C.items=B:Object.assign(C,B)}else C.items=g;(w=C.items)!=null&&w.length&&(C.items=C.items.filter(B=>!B.hidden),E(C))})}function V(n){s.expandedKeys=n}function Te(){return s.expandedKeys}function ne(n){s.selectedKeys=n}function _e(){return s.selectedKeys}function ae(n){s.checkedKeys=n}function Be(){return s.checkedKeys}function le(n){s.checkedKeys=n?m():[]}function G(n){s.expandedKeys=n?b():[]}function Fe(n){s.checkStrictly=n}$(()=>e.searchValue,n=>{n!==p.searchText&&(p.searchText=n)},{immediate:!0}),$(()=>e.treeData,n=>{n&&J(p.searchText)});function J(n){if(n!==p.searchText&&(p.searchText=n),k("update:searchValue",n),!n){p.startSearch=!1;return}const{filterFn:f,checkable:x,expandOnSearch:g,checkOnSearch:S,selectedOnSearch:C}=u(e);p.startSearch=!0;const{title:w,key:B}=u(v),D=[];if(p.searchData=at(u(K),F=>{var X,Y;const N=f?f(n,F,u(v)):(Y=(X=F[w])==null?void 0:X.includes(n))!=null?Y:!1;return N&&D.push(F[B]),N},u(v)),g){const F=lt(p.searchData).map(N=>N[B]);F&&F.length&&V(F)}S&&x&&D.length&&ae(D),C&&D.length&&ne(D)}function we(n,f){if(!(!e.clickRowToExpand||!f||f.length===0))if(!s.expandedKeys.includes(n))V([...s.expandedKeys,n]);else{const x=[...s.expandedKeys],g=x.findIndex(S=>S===n);g!==-1&&x.splice(g,1),V(x)}}H(()=>{K.value=e.treeData}),Je(()=>{const n=parseInt(e.defaultExpandLevel);n>0?s.expandedKeys=r(n):e.defaultExpandAll&&G(!0)}),H(()=>{s.expandedKeys=e.expandedKeys}),H(()=>{s.selectedKeys=e.selectedKeys}),H(()=>{s.checkedKeys=e.checkedKeys}),$(()=>e.value,()=>{s.checkedKeys=R(e.value||[])},{immediate:!0}),$(()=>s.checkedKeys,()=>{const n=R(s.checkedKeys);k("update:value",n),k("change",n)}),H(()=>{s.checkStrictly=e.checkStrictly});const De={setExpandedKeys:V,getExpandedKeys:Te,setSelectedKeys:ne,getSelectedKeys:_e,setCheckedKeys:ae,getCheckedKeys:Be,insertNodeByKey:d,insertNodesByKey:o,deleteNodeByKey:i,updateNodeByKey:y,getSelectedNode:ve,checkAll:le,expandAll:G,filterByLevel:n=>{s.expandedKeys=r(n)},setSearchValue:n=>{J(n)},getSearchValue:()=>p.searchText};function Ne(n){const{actionList:f}=e;if(!(!f||f.length===0))return f.map((x,g)=>{var C;let S=!0;return W(x.show)?S=(C=x.show)==null?void 0:C.call(x,n):ge(x.show)&&(S=x.show),S?L("span",{key:g,class:A("action")},[x.render(n)]):null})}const Ie=M(()=>{const n=ee(t.value);return ct(n,(f,x)=>{var ce;const g=p.searchText,{highlight:S}=u(e),{title:C,key:w,children:B}=u(v),D=Ee(f,f.icon),F=At(f,C),N=g?F.indexOf(g):-1,X=p.startSearch&&!it(g)&&S&&N!==-1,Y=`color: ${ge(S)?"#f50":S}`,Me=X?L("span",{class:(ce=u(h))!=null&&ce.blockNode?`${A("content")}`:""},[L("span",null,[F.substr(0,N)]),L("span",{style:Y},[g]),L("span",null,[F.substr(N+g.length)])]):F;return f[C]=L("span",{class:`${A("title")} pl-2`,onClick:we.bind(null,f[w],f[B])},[l!=null&&l.title?ut(l,"title",f):L(Q,null,[D&&L(wt,{icon:D},null),Me,L("span",{class:A("actions")},[Ne(f)])])]),f}),n});return _(De),()=>{let n;const{title:f,helpMessage:x,toolbar:g,search:S,checkable:C}=e,w=f||g||S||l.headerTitle,B={height:"calc(100% - 38px)"};return L("div",{class:[A(),"h-full",a.class]},[w&&L(Ft,{checkable:C,checkAll:le,expandAll:G,title:f,search:S,toolbar:g,helpMessage:x,onStrictlyChange:Fe,onSearch:J,searchText:p.searchText},Nt(n=st(l))?n:{default:()=>[n]}),L(xt,{wrapperClassName:u(e.treeWrapperClassName),spinning:u(e.loading),tip:"加载中..."},{default:()=>[ye(L(rt,{style:B},{default:()=>[L(Ct,qe(u(h),{showIcon:!1,treeData:Ie.value}),null)]}),[[pe,!u(c)]]),ye(L(Se,{image:Se.PRESENTED_IMAGE_SIMPLE,class:"!mt-4"},null),[[pe,u(c)]])]})])}}});export{mt as _};