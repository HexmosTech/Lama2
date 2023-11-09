/******/ (function(modules) { // webpackBootstrap
/******/ 	// The module cache
/******/ 	var installedModules = {};
/******/
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/
/******/ 		// Check if module is in cache
/******/ 		if(installedModules[moduleId]) {
/******/ 			return installedModules[moduleId].exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = installedModules[moduleId] = {
/******/ 			i: moduleId,
/******/ 			l: false,
/******/ 			exports: {}
/******/ 		};
/******/
/******/ 		// Execute the module function
/******/ 		modules[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/
/******/ 		// Flag the module as loaded
/******/ 		module.l = true;
/******/
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/
/******/
/******/ 	// expose the modules object (__webpack_modules__)
/******/ 	__webpack_require__.m = modules;
/******/
/******/ 	// expose the module cache
/******/ 	__webpack_require__.c = installedModules;
/******/
/******/ 	// define getter function for harmony exports
/******/ 	__webpack_require__.d = function(exports, name, getter) {
/******/ 		if(!__webpack_require__.o(exports, name)) {
/******/ 			Object.defineProperty(exports, name, {
/******/ 				configurable: false,
/******/ 				enumerable: true,
/******/ 				get: getter
/******/ 			});
/******/ 		}
/******/ 	};
/******/
/******/ 	// getDefaultExport function for compatibility with non-harmony modules
/******/ 	__webpack_require__.n = function(module) {
/******/ 		var getter = module && module.__esModule ?
/******/ 			function getDefault() { return module['default']; } :
/******/ 			function getModuleExports() { return module; };
/******/ 		__webpack_require__.d(getter, 'a', getter);
/******/ 		return getter;
/******/ 	};
/******/
/******/ 	// Object.prototype.hasOwnProperty.call
/******/ 	__webpack_require__.o = function(object, property) { return Object.prototype.hasOwnProperty.call(object, property); };
/******/
/******/ 	// __webpack_public_path__
/******/ 	__webpack_require__.p = "";
/******/
/******/ 	// Load entry module and return exports
/******/ 	return __webpack_require__(__webpack_require__.s = 0);
/******/ })
/************************************************************************/
/******/ ([
/* 0 */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__services__ = __webpack_require__(1);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__views_lb__ = __webpack_require__(2);
function _typeof(o) { "@babel/helpers - typeof"; return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (o) { return typeof o; } : function (o) { return o && "function" == typeof Symbol && o.constructor === Symbol && o !== Symbol.prototype ? "symbol" : typeof o; }, _typeof(o); }
function _regeneratorRuntime() { "use strict"; /*! regenerator-runtime -- Copyright (c) 2014-present, Facebook, Inc. -- license (MIT): https://github.com/facebook/regenerator/blob/main/LICENSE */ _regeneratorRuntime = function _regeneratorRuntime() { return e; }; var t, e = {}, r = Object.prototype, n = r.hasOwnProperty, o = Object.defineProperty || function (t, e, r) { t[e] = r.value; }, i = "function" == typeof Symbol ? Symbol : {}, a = i.iterator || "@@iterator", c = i.asyncIterator || "@@asyncIterator", u = i.toStringTag || "@@toStringTag"; function define(t, e, r) { return Object.defineProperty(t, e, { value: r, enumerable: !0, configurable: !0, writable: !0 }), t[e]; } try { define({}, ""); } catch (t) { define = function define(t, e, r) { return t[e] = r; }; } function wrap(t, e, r, n) { var i = e && e.prototype instanceof Generator ? e : Generator, a = Object.create(i.prototype), c = new Context(n || []); return o(a, "_invoke", { value: makeInvokeMethod(t, r, c) }), a; } function tryCatch(t, e, r) { try { return { type: "normal", arg: t.call(e, r) }; } catch (t) { return { type: "throw", arg: t }; } } e.wrap = wrap; var h = "suspendedStart", l = "suspendedYield", f = "executing", s = "completed", y = {}; function Generator() {} function GeneratorFunction() {} function GeneratorFunctionPrototype() {} var p = {}; define(p, a, function () { return this; }); var d = Object.getPrototypeOf, v = d && d(d(values([]))); v && v !== r && n.call(v, a) && (p = v); var g = GeneratorFunctionPrototype.prototype = Generator.prototype = Object.create(p); function defineIteratorMethods(t) { ["next", "throw", "return"].forEach(function (e) { define(t, e, function (t) { return this._invoke(e, t); }); }); } function AsyncIterator(t, e) { function invoke(r, o, i, a) { var c = tryCatch(t[r], t, o); if ("throw" !== c.type) { var u = c.arg, h = u.value; return h && "object" == _typeof(h) && n.call(h, "__await") ? e.resolve(h.__await).then(function (t) { invoke("next", t, i, a); }, function (t) { invoke("throw", t, i, a); }) : e.resolve(h).then(function (t) { u.value = t, i(u); }, function (t) { return invoke("throw", t, i, a); }); } a(c.arg); } var r; o(this, "_invoke", { value: function value(t, n) { function callInvokeWithMethodAndArg() { return new e(function (e, r) { invoke(t, n, e, r); }); } return r = r ? r.then(callInvokeWithMethodAndArg, callInvokeWithMethodAndArg) : callInvokeWithMethodAndArg(); } }); } function makeInvokeMethod(e, r, n) { var o = h; return function (i, a) { if (o === f) throw new Error("Generator is already running"); if (o === s) { if ("throw" === i) throw a; return { value: t, done: !0 }; } for (n.method = i, n.arg = a;;) { var c = n.delegate; if (c) { var u = maybeInvokeDelegate(c, n); if (u) { if (u === y) continue; return u; } } if ("next" === n.method) n.sent = n._sent = n.arg;else if ("throw" === n.method) { if (o === h) throw o = s, n.arg; n.dispatchException(n.arg); } else "return" === n.method && n.abrupt("return", n.arg); o = f; var p = tryCatch(e, r, n); if ("normal" === p.type) { if (o = n.done ? s : l, p.arg === y) continue; return { value: p.arg, done: n.done }; } "throw" === p.type && (o = s, n.method = "throw", n.arg = p.arg); } }; } function maybeInvokeDelegate(e, r) { var n = r.method, o = e.iterator[n]; if (o === t) return r.delegate = null, "throw" === n && e.iterator["return"] && (r.method = "return", r.arg = t, maybeInvokeDelegate(e, r), "throw" === r.method) || "return" !== n && (r.method = "throw", r.arg = new TypeError("The iterator does not provide a '" + n + "' method")), y; var i = tryCatch(o, e.iterator, r.arg); if ("throw" === i.type) return r.method = "throw", r.arg = i.arg, r.delegate = null, y; var a = i.arg; return a ? a.done ? (r[e.resultName] = a.value, r.next = e.nextLoc, "return" !== r.method && (r.method = "next", r.arg = t), r.delegate = null, y) : a : (r.method = "throw", r.arg = new TypeError("iterator result is not an object"), r.delegate = null, y); } function pushTryEntry(t) { var e = { tryLoc: t[0] }; 1 in t && (e.catchLoc = t[1]), 2 in t && (e.finallyLoc = t[2], e.afterLoc = t[3]), this.tryEntries.push(e); } function resetTryEntry(t) { var e = t.completion || {}; e.type = "normal", delete e.arg, t.completion = e; } function Context(t) { this.tryEntries = [{ tryLoc: "root" }], t.forEach(pushTryEntry, this), this.reset(!0); } function values(e) { if (e || "" === e) { var r = e[a]; if (r) return r.call(e); if ("function" == typeof e.next) return e; if (!isNaN(e.length)) { var o = -1, i = function next() { for (; ++o < e.length;) if (n.call(e, o)) return next.value = e[o], next.done = !1, next; return next.value = t, next.done = !0, next; }; return i.next = i; } } throw new TypeError(_typeof(e) + " is not iterable"); } return GeneratorFunction.prototype = GeneratorFunctionPrototype, o(g, "constructor", { value: GeneratorFunctionPrototype, configurable: !0 }), o(GeneratorFunctionPrototype, "constructor", { value: GeneratorFunction, configurable: !0 }), GeneratorFunction.displayName = define(GeneratorFunctionPrototype, u, "GeneratorFunction"), e.isGeneratorFunction = function (t) { var e = "function" == typeof t && t.constructor; return !!e && (e === GeneratorFunction || "GeneratorFunction" === (e.displayName || e.name)); }, e.mark = function (t) { return Object.setPrototypeOf ? Object.setPrototypeOf(t, GeneratorFunctionPrototype) : (t.__proto__ = GeneratorFunctionPrototype, define(t, u, "GeneratorFunction")), t.prototype = Object.create(g), t; }, e.awrap = function (t) { return { __await: t }; }, defineIteratorMethods(AsyncIterator.prototype), define(AsyncIterator.prototype, c, function () { return this; }), e.AsyncIterator = AsyncIterator, e.async = function (t, r, n, o, i) { void 0 === i && (i = Promise); var a = new AsyncIterator(wrap(t, r, n, o), i); return e.isGeneratorFunction(r) ? a : a.next().then(function (t) { return t.done ? t.value : a.next(); }); }, defineIteratorMethods(g), define(g, u, "Generator"), define(g, a, function () { return this; }), define(g, "toString", function () { return "[object Generator]"; }), e.keys = function (t) { var e = Object(t), r = []; for (var n in e) r.push(n); return r.reverse(), function next() { for (; r.length;) { var t = r.pop(); if (t in e) return next.value = t, next.done = !1, next; } return next.done = !0, next; }; }, e.values = values, Context.prototype = { constructor: Context, reset: function reset(e) { if (this.prev = 0, this.next = 0, this.sent = this._sent = t, this.done = !1, this.delegate = null, this.method = "next", this.arg = t, this.tryEntries.forEach(resetTryEntry), !e) for (var r in this) "t" === r.charAt(0) && n.call(this, r) && !isNaN(+r.slice(1)) && (this[r] = t); }, stop: function stop() { this.done = !0; var t = this.tryEntries[0].completion; if ("throw" === t.type) throw t.arg; return this.rval; }, dispatchException: function dispatchException(e) { if (this.done) throw e; var r = this; function handle(n, o) { return a.type = "throw", a.arg = e, r.next = n, o && (r.method = "next", r.arg = t), !!o; } for (var o = this.tryEntries.length - 1; o >= 0; --o) { var i = this.tryEntries[o], a = i.completion; if ("root" === i.tryLoc) return handle("end"); if (i.tryLoc <= this.prev) { var c = n.call(i, "catchLoc"), u = n.call(i, "finallyLoc"); if (c && u) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } else if (c) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); } else { if (!u) throw new Error("try statement without catch or finally"); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } } } }, abrupt: function abrupt(t, e) { for (var r = this.tryEntries.length - 1; r >= 0; --r) { var o = this.tryEntries[r]; if (o.tryLoc <= this.prev && n.call(o, "finallyLoc") && this.prev < o.finallyLoc) { var i = o; break; } } i && ("break" === t || "continue" === t) && i.tryLoc <= e && e <= i.finallyLoc && (i = null); var a = i ? i.completion : {}; return a.type = t, a.arg = e, i ? (this.method = "next", this.next = i.finallyLoc, y) : this.complete(a); }, complete: function complete(t, e) { if ("throw" === t.type) throw t.arg; return "break" === t.type || "continue" === t.type ? this.next = t.arg : "return" === t.type ? (this.rval = this.arg = t.arg, this.method = "return", this.next = "end") : "normal" === t.type && e && (this.next = e), y; }, finish: function finish(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.finallyLoc === t) return this.complete(r.completion, r.afterLoc), resetTryEntry(r), y; } }, "catch": function _catch(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.tryLoc === t) { var n = r.completion; if ("throw" === n.type) { var o = n.arg; resetTryEntry(r); } return o; } } throw new Error("illegal catch attempt"); }, delegateYield: function delegateYield(e, r, n) { return this.delegate = { iterator: values(e), resultName: r, nextLoc: n }, "next" === this.method && (this.arg = t), y; } }, e; }
function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { Promise.resolve(value).then(_next, _throw); } }
function _asyncToGenerator(fn) { return function () { var self = this, args = arguments; return new Promise(function (resolve, reject) { var gen = fn.apply(self, args); function _next(value) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value); } function _throw(err) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err); } _next(undefined); }); }; }

// import { show } from './views/message'

var supportedAPI = ['init', 'lamabutton']; // enlist all methods supported by API (e.g. `mw('event', 'user-login');`)
var go;
/**
    The main entry of the application
    */
function app(window) {
  console.log('JS-Widget starting');

  // set default configurations
  var configurations = {
    someDefaultConfiguration: false
  };

  // all methods that were called till now and stored in queue
  // needs to be called now 
  var globalObject = window[window['JS-Widget']];
  var queue = globalObject.q;
  if (queue) {
    for (var i = 0; i < queue.length; i++) {
      if (queue[i][0].toLowerCase() == 'init') {
        configurations = extendObject(configurations, queue[i][1]);
        console.log('JS-Widget started', configurations);
      } else apiHandler(queue[i][0], queue[i][1]);
    }
  }

  // override temporary (until the app loaded) handler
  // for widget's API calls
  globalObject = apiHandler;
  globalObject.configurations = configurations;
}

/**
    Method that handles all API calls
    */
function apiHandler(_x, _x2) {
  return _apiHandler.apply(this, arguments);
}
function _apiHandler() {
  _apiHandler = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee(api, params) {
    return _regeneratorRuntime().wrap(function _callee$(_context) {
      while (1) switch (_context.prev = _context.next) {
        case 0:
          if (api) {
            _context.next = 2;
            break;
          }
          throw Error('API method required');
        case 2:
          api = api.toLowerCase();
          if (!(supportedAPI.indexOf(api) === -1)) {
            _context.next = 5;
            break;
          }
          throw Error("Method ".concat(api, " is not supported"));
        case 5:
          console.log("Handling API call ".concat(api), params);
          _context.t0 = api;
          _context.next = _context.t0 === 'lamabutton' ? 9 : 14;
          break;
        case 9:
          _context.next = 11;
          return loadWasm();
        case 11:
          go = _context.sent;
          // Ensure Wasm is loaded and Go instance is created
          //     if (!go) return; // Exit if there was an error loading Wasm
          // }
          Object(__WEBPACK_IMPORTED_MODULE_1__views_lb__["a" /* addButtonToTextarea */])(params, go);
          return _context.abrupt("break", 15);
        case 14:
          console.warn("No handler defined for ".concat(api));
        case 15:
        case "end":
          return _context.stop();
      }
    }, _callee);
  }));
  return _apiHandler.apply(this, arguments);
}
function extendObject(a, b) {
  for (var key in b) if (b.hasOwnProperty(key)) a[key] = b[key];
  return a;
}
app(window);
function loadWasm() {
  return _loadWasm.apply(this, arguments);
}
function _loadWasm() {
  _loadWasm = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee2() {
    var response, buffer, _go, result;
    return _regeneratorRuntime().wrap(function _callee2$(_context2) {
      while (1) switch (_context2.prev = _context2.next) {
        case 0:
          _context2.prev = 0;
          _context2.next = 3;
          return fetch("main.wasm");
        case 3:
          response = _context2.sent;
          _context2.next = 6;
          return response.arrayBuffer();
        case 6:
          buffer = _context2.sent;
          _go = new Go();
          _context2.next = 10;
          return WebAssembly.instantiate(buffer, _go.importObject);
        case 10:
          result = _context2.sent;
          _go.run(result.instance);
          console.log("Wasm loaded and Go instance initialized");
          return _context2.abrupt("return", _go);
        case 16:
          _context2.prev = 16;
          _context2.t0 = _context2["catch"](0);
          console.error("Error loading Wasm:", _context2.t0);
          return _context2.abrupt("return", null);
        case 20:
        case "end":
          return _context2.stop();
      }
    }, _callee2, null, [[0, 16]]);
  }));
  return _loadWasm.apply(this, arguments);
}

/***/ }),
/* 1 */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* unused harmony export ping */
function ping() {
  return 'pong';
}

/***/ }),
/* 2 */
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
/* harmony export (immutable) */ __webpack_exports__["a"] = addButtonToTextarea;
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__lb_css__ = __webpack_require__(3);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__lb_css___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__lb_css__);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__wasm_exec__ = __webpack_require__(8);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__wasm_exec___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__wasm_exec__);
function _typeof(o) { "@babel/helpers - typeof"; return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (o) { return typeof o; } : function (o) { return o && "function" == typeof Symbol && o.constructor === Symbol && o !== Symbol.prototype ? "symbol" : typeof o; }, _typeof(o); }
function _regeneratorRuntime() { "use strict"; /*! regenerator-runtime -- Copyright (c) 2014-present, Facebook, Inc. -- license (MIT): https://github.com/facebook/regenerator/blob/main/LICENSE */ _regeneratorRuntime = function _regeneratorRuntime() { return e; }; var t, e = {}, r = Object.prototype, n = r.hasOwnProperty, o = Object.defineProperty || function (t, e, r) { t[e] = r.value; }, i = "function" == typeof Symbol ? Symbol : {}, a = i.iterator || "@@iterator", c = i.asyncIterator || "@@asyncIterator", u = i.toStringTag || "@@toStringTag"; function define(t, e, r) { return Object.defineProperty(t, e, { value: r, enumerable: !0, configurable: !0, writable: !0 }), t[e]; } try { define({}, ""); } catch (t) { define = function define(t, e, r) { return t[e] = r; }; } function wrap(t, e, r, n) { var i = e && e.prototype instanceof Generator ? e : Generator, a = Object.create(i.prototype), c = new Context(n || []); return o(a, "_invoke", { value: makeInvokeMethod(t, r, c) }), a; } function tryCatch(t, e, r) { try { return { type: "normal", arg: t.call(e, r) }; } catch (t) { return { type: "throw", arg: t }; } } e.wrap = wrap; var h = "suspendedStart", l = "suspendedYield", f = "executing", s = "completed", y = {}; function Generator() {} function GeneratorFunction() {} function GeneratorFunctionPrototype() {} var p = {}; define(p, a, function () { return this; }); var d = Object.getPrototypeOf, v = d && d(d(values([]))); v && v !== r && n.call(v, a) && (p = v); var g = GeneratorFunctionPrototype.prototype = Generator.prototype = Object.create(p); function defineIteratorMethods(t) { ["next", "throw", "return"].forEach(function (e) { define(t, e, function (t) { return this._invoke(e, t); }); }); } function AsyncIterator(t, e) { function invoke(r, o, i, a) { var c = tryCatch(t[r], t, o); if ("throw" !== c.type) { var u = c.arg, h = u.value; return h && "object" == _typeof(h) && n.call(h, "__await") ? e.resolve(h.__await).then(function (t) { invoke("next", t, i, a); }, function (t) { invoke("throw", t, i, a); }) : e.resolve(h).then(function (t) { u.value = t, i(u); }, function (t) { return invoke("throw", t, i, a); }); } a(c.arg); } var r; o(this, "_invoke", { value: function value(t, n) { function callInvokeWithMethodAndArg() { return new e(function (e, r) { invoke(t, n, e, r); }); } return r = r ? r.then(callInvokeWithMethodAndArg, callInvokeWithMethodAndArg) : callInvokeWithMethodAndArg(); } }); } function makeInvokeMethod(e, r, n) { var o = h; return function (i, a) { if (o === f) throw new Error("Generator is already running"); if (o === s) { if ("throw" === i) throw a; return { value: t, done: !0 }; } for (n.method = i, n.arg = a;;) { var c = n.delegate; if (c) { var u = maybeInvokeDelegate(c, n); if (u) { if (u === y) continue; return u; } } if ("next" === n.method) n.sent = n._sent = n.arg;else if ("throw" === n.method) { if (o === h) throw o = s, n.arg; n.dispatchException(n.arg); } else "return" === n.method && n.abrupt("return", n.arg); o = f; var p = tryCatch(e, r, n); if ("normal" === p.type) { if (o = n.done ? s : l, p.arg === y) continue; return { value: p.arg, done: n.done }; } "throw" === p.type && (o = s, n.method = "throw", n.arg = p.arg); } }; } function maybeInvokeDelegate(e, r) { var n = r.method, o = e.iterator[n]; if (o === t) return r.delegate = null, "throw" === n && e.iterator["return"] && (r.method = "return", r.arg = t, maybeInvokeDelegate(e, r), "throw" === r.method) || "return" !== n && (r.method = "throw", r.arg = new TypeError("The iterator does not provide a '" + n + "' method")), y; var i = tryCatch(o, e.iterator, r.arg); if ("throw" === i.type) return r.method = "throw", r.arg = i.arg, r.delegate = null, y; var a = i.arg; return a ? a.done ? (r[e.resultName] = a.value, r.next = e.nextLoc, "return" !== r.method && (r.method = "next", r.arg = t), r.delegate = null, y) : a : (r.method = "throw", r.arg = new TypeError("iterator result is not an object"), r.delegate = null, y); } function pushTryEntry(t) { var e = { tryLoc: t[0] }; 1 in t && (e.catchLoc = t[1]), 2 in t && (e.finallyLoc = t[2], e.afterLoc = t[3]), this.tryEntries.push(e); } function resetTryEntry(t) { var e = t.completion || {}; e.type = "normal", delete e.arg, t.completion = e; } function Context(t) { this.tryEntries = [{ tryLoc: "root" }], t.forEach(pushTryEntry, this), this.reset(!0); } function values(e) { if (e || "" === e) { var r = e[a]; if (r) return r.call(e); if ("function" == typeof e.next) return e; if (!isNaN(e.length)) { var o = -1, i = function next() { for (; ++o < e.length;) if (n.call(e, o)) return next.value = e[o], next.done = !1, next; return next.value = t, next.done = !0, next; }; return i.next = i; } } throw new TypeError(_typeof(e) + " is not iterable"); } return GeneratorFunction.prototype = GeneratorFunctionPrototype, o(g, "constructor", { value: GeneratorFunctionPrototype, configurable: !0 }), o(GeneratorFunctionPrototype, "constructor", { value: GeneratorFunction, configurable: !0 }), GeneratorFunction.displayName = define(GeneratorFunctionPrototype, u, "GeneratorFunction"), e.isGeneratorFunction = function (t) { var e = "function" == typeof t && t.constructor; return !!e && (e === GeneratorFunction || "GeneratorFunction" === (e.displayName || e.name)); }, e.mark = function (t) { return Object.setPrototypeOf ? Object.setPrototypeOf(t, GeneratorFunctionPrototype) : (t.__proto__ = GeneratorFunctionPrototype, define(t, u, "GeneratorFunction")), t.prototype = Object.create(g), t; }, e.awrap = function (t) { return { __await: t }; }, defineIteratorMethods(AsyncIterator.prototype), define(AsyncIterator.prototype, c, function () { return this; }), e.AsyncIterator = AsyncIterator, e.async = function (t, r, n, o, i) { void 0 === i && (i = Promise); var a = new AsyncIterator(wrap(t, r, n, o), i); return e.isGeneratorFunction(r) ? a : a.next().then(function (t) { return t.done ? t.value : a.next(); }); }, defineIteratorMethods(g), define(g, u, "Generator"), define(g, a, function () { return this; }), define(g, "toString", function () { return "[object Generator]"; }), e.keys = function (t) { var e = Object(t), r = []; for (var n in e) r.push(n); return r.reverse(), function next() { for (; r.length;) { var t = r.pop(); if (t in e) return next.value = t, next.done = !1, next; } return next.done = !0, next; }; }, e.values = values, Context.prototype = { constructor: Context, reset: function reset(e) { if (this.prev = 0, this.next = 0, this.sent = this._sent = t, this.done = !1, this.delegate = null, this.method = "next", this.arg = t, this.tryEntries.forEach(resetTryEntry), !e) for (var r in this) "t" === r.charAt(0) && n.call(this, r) && !isNaN(+r.slice(1)) && (this[r] = t); }, stop: function stop() { this.done = !0; var t = this.tryEntries[0].completion; if ("throw" === t.type) throw t.arg; return this.rval; }, dispatchException: function dispatchException(e) { if (this.done) throw e; var r = this; function handle(n, o) { return a.type = "throw", a.arg = e, r.next = n, o && (r.method = "next", r.arg = t), !!o; } for (var o = this.tryEntries.length - 1; o >= 0; --o) { var i = this.tryEntries[o], a = i.completion; if ("root" === i.tryLoc) return handle("end"); if (i.tryLoc <= this.prev) { var c = n.call(i, "catchLoc"), u = n.call(i, "finallyLoc"); if (c && u) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } else if (c) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); } else { if (!u) throw new Error("try statement without catch or finally"); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } } } }, abrupt: function abrupt(t, e) { for (var r = this.tryEntries.length - 1; r >= 0; --r) { var o = this.tryEntries[r]; if (o.tryLoc <= this.prev && n.call(o, "finallyLoc") && this.prev < o.finallyLoc) { var i = o; break; } } i && ("break" === t || "continue" === t) && i.tryLoc <= e && e <= i.finallyLoc && (i = null); var a = i ? i.completion : {}; return a.type = t, a.arg = e, i ? (this.method = "next", this.next = i.finallyLoc, y) : this.complete(a); }, complete: function complete(t, e) { if ("throw" === t.type) throw t.arg; return "break" === t.type || "continue" === t.type ? this.next = t.arg : "return" === t.type ? (this.rval = this.arg = t.arg, this.method = "return", this.next = "end") : "normal" === t.type && e && (this.next = e), y; }, finish: function finish(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.finallyLoc === t) return this.complete(r.completion, r.afterLoc), resetTryEntry(r), y; } }, "catch": function _catch(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.tryLoc === t) { var n = r.completion; if ("throw" === n.type) { var o = n.arg; resetTryEntry(r); } return o; } } throw new Error("illegal catch attempt"); }, delegateYield: function delegateYield(e, r, n) { return this.delegate = { iterator: values(e), resultName: r, nextLoc: n }, "next" === this.method && (this.arg = t), y; } }, e; }
function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { Promise.resolve(value).then(_next, _throw); } }
function _asyncToGenerator(fn) { return function () { var self = this, args = arguments; return new Promise(function (resolve, reject) { var gen = fn.apply(self, args); function _next(value) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value); } function _throw(err) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err); } _next(undefined); }); }; }


function addButtonToTextarea(className, go) {
  if (!go) {
    console.log("wasm file not initialized");
    return;
  }
  // const divElements = document.querySelectorAll(`div.${className}`);
  // Select all elements with the class name 'my-button-container'
  var divElements = document.querySelectorAll(".".concat(className));
  divElements.forEach(function (divElement) {
    var buttonImage = document.createElement('img');
    buttonImage.src = 'arrowl2.png';
    buttonImage.style.width = '30px';
    var index = Math.random().toString(36).substr(2, 10);
    var buttonElement = document.createElement('button');
    buttonElement.appendChild(buttonImage);
    buttonElement.addEventListener('click', function () {
      return addTextToDiv(divElement, index);
    });
    buttonElement.style.all = 'unset';
    buttonElement.style.position = 'relative';
    buttonElement.style.bottom = '150px';
    buttonElement.style.right = '50px';

    // Add hover effect
    buttonElement.style.transition = 'background-color 0.3s';
    buttonElement.style.cursor = 'pointer';
    var highlightColor = '#00A1E1';
    buttonElement.addEventListener('mouseenter', function () {
      buttonElement.style.backgroundColor = highlightColor; // Change the background color on hover
    });

    buttonElement.addEventListener('mouseleave', function () {
      buttonElement.style.backgroundColor = ''; // Reset the background color on mouse leave
    });

    buttonElement.addEventListener('mousedown', function () {
      buttonElement.style.backgroundColor = highlightColor; // Change the background color when clicked
    });

    buttonElement.addEventListener('mouseup', function () {
      buttonElement.style.backgroundColor = highlightColor; // Restore the hover color on release
    });

    divElement.appendChild(buttonElement);
    var outputDiv = document.createElement('div');
    outputDiv["class"] = 'outputField';
    outputDiv.id = "outputField_".concat(index);
    outputDiv.style.fontSize = '16px';
    outputDiv.style.marginTop = '10px';
    outputDiv.style.width = '500px';
    outputDiv.style.height = '300px';
    outputDiv.style.color = 'white';
    outputDiv.style.backgroundColor = 'black';
    outputDiv.style.border = '1px solid #555';
    outputDiv.style.padding = '10px';
    outputDiv.style.fontFamily = 'monospace';
    outputDiv.style.resize = 'vertical';
    outputDiv.style.overflowY = 'auto';
    outputDiv.style.whiteSpace = 'pre-wrap';
    var logoImage = document.createElement('img');
    logoImage.style.position = 'relative';
    logoImage.style.bottom = '10px';
    logoImage.style.right = '80px';
    logoImage.src = 'hexmos_logo.jpeg';
    logoImage.style.width = '30px';
    logoImage.style.borderRadius = '50%';
    logoImage.addEventListener('click', function () {
      window.location.href = 'https://hexmos.com';
    });
    divElement.appendChild(logoImage);
    divElement.appendChild(outputDiv);
  });
}
function addTextToDiv(_x, _x2) {
  return _addTextToDiv.apply(this, arguments);
}
function _addTextToDiv() {
  _addTextToDiv = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee(divElement, index) {
    var textArea;
    return _regeneratorRuntime().wrap(function _callee$(_context) {
      while (1) switch (_context.prev = _context.next) {
        case 0:
          textArea = divElement.querySelector('textarea');
          if (textArea) {
            console.log("Text inside the text area:", textArea.value);
            makeLamaRequest(textArea.value, divElement, index);
          } else {
            console.log("No text area found. Logging data from divElement:", divElement.textContent);
            makeLamaRequest(divElement.textContent, divElement, index);
          }
        case 2:
        case "end":
          return _context.stop();
      }
    }, _callee);
  }));
  return _addTextToDiv.apply(this, arguments);
}
function makeLamaRequest(_x3, _x4, _x5) {
  return _makeLamaRequest.apply(this, arguments);
}
function _makeLamaRequest() {
  _makeLamaRequest = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee2(command, divElement, index) {
    var response, jsonObject, formattedBody, formattedJson, outputField;
    return _regeneratorRuntime().wrap(function _callee2$(_context2) {
      while (1) switch (_context2.prev = _context2.next) {
        case 0:
          _context2.prev = 0;
          if (go) {
            _context2.next = 4;
            break;
          }
          console.error("Wasm is not loaded. Cannot make a request.");
          return _context2.abrupt("return");
        case 4:
          _context2.next = 6;
          return goWebRequestFunc(command);
        case 6:
          response = _context2.sent;
          console.log("response Data:", response);
          jsonObject = JSON.parse(response);
          formattedBody = JSON.parse(jsonObject.Body);
          formattedJson = JSON.stringify(formattedBody, null, 2);
          console.log("formattedJson:", formattedJson);
          outputField = divElement.querySelector("#outputField_".concat(index));
          console.log(outputField);
          outputField.textContent = formattedJson;
          console.log("outputField.textContent:", outputField.textContent);
          _context2.next = 21;
          break;
        case 18:
          _context2.prev = 18;
          _context2.t0 = _context2["catch"](0);
          console.error("Error making Lama request:", _context2.t0);
        case 21:
        case "end":
          return _context2.stop();
      }
    }, _callee2, null, [[0, 18]]);
  }));
  return _makeLamaRequest.apply(this, arguments);
}

/***/ }),
/* 3 */
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(4);
if(typeof content === 'string') content = [[module.i, content, '']];
// Prepare cssTransformation
var transform;

var options = {"hmr":true}
options.transform = transform
// add the styles to the DOM
var update = __webpack_require__(6)(content, options);
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../node_modules/css-loader/dist/cjs.js!./lb.css", function() {
			var newContent = require("!!../../node_modules/css-loader/dist/cjs.js!./lb.css");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),
/* 4 */
/***/ (function(module, exports, __webpack_require__) {

// Imports
var ___CSS_LOADER_API_IMPORT___ = __webpack_require__(5);
exports = ___CSS_LOADER_API_IMPORT___(false);
// Module
exports.push([module.i, "/* lamabutton.css */\n\n/* Style for the button element */\n.js-widget-lamabutton {\n    background-color: green; /* Set the background color to green */\n    color: white; /* Set the text color to white */\n    border: none; /* Remove borders */\n    padding: 5px 10px; /* Adjust padding for a smaller button */\n    cursor: pointer; /* Add a pointer cursor on hover */\n    position: absolute;\n}\n\n/* Hover effect when the button is hovered over */\n.js-widget-lamabutton:hover {\n    background-color: darkgreen; /* Darken the background color on hover */\n}\n", ""]);
// Exports
module.exports = exports;


/***/ }),
/* 5 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";


/*
  MIT License http://www.opensource.org/licenses/mit-license.php
  Author Tobias Koppers @sokra
*/
// css base code, injected by the css-loader
// eslint-disable-next-line func-names
module.exports = function (useSourceMap) {
  var list = []; // return the list of modules as css string

  list.toString = function toString() {
    return this.map(function (item) {
      var content = cssWithMappingToString(item, useSourceMap);

      if (item[2]) {
        return "@media ".concat(item[2], " {").concat(content, "}");
      }

      return content;
    }).join('');
  }; // import a list of modules into the list
  // eslint-disable-next-line func-names


  list.i = function (modules, mediaQuery, dedupe) {
    if (typeof modules === 'string') {
      // eslint-disable-next-line no-param-reassign
      modules = [[null, modules, '']];
    }

    var alreadyImportedModules = {};

    if (dedupe) {
      for (var i = 0; i < this.length; i++) {
        // eslint-disable-next-line prefer-destructuring
        var id = this[i][0];

        if (id != null) {
          alreadyImportedModules[id] = true;
        }
      }
    }

    for (var _i = 0; _i < modules.length; _i++) {
      var item = [].concat(modules[_i]);

      if (dedupe && alreadyImportedModules[item[0]]) {
        // eslint-disable-next-line no-continue
        continue;
      }

      if (mediaQuery) {
        if (!item[2]) {
          item[2] = mediaQuery;
        } else {
          item[2] = "".concat(mediaQuery, " and ").concat(item[2]);
        }
      }

      list.push(item);
    }
  };

  return list;
};

function cssWithMappingToString(item, useSourceMap) {
  var content = item[1] || ''; // eslint-disable-next-line prefer-destructuring

  var cssMapping = item[3];

  if (!cssMapping) {
    return content;
  }

  if (useSourceMap && typeof btoa === 'function') {
    var sourceMapping = toComment(cssMapping);
    var sourceURLs = cssMapping.sources.map(function (source) {
      return "/*# sourceURL=".concat(cssMapping.sourceRoot || '').concat(source, " */");
    });
    return [content].concat(sourceURLs).concat([sourceMapping]).join('\n');
  }

  return [content].join('\n');
} // Adapted from convert-source-map (MIT)


function toComment(sourceMap) {
  // eslint-disable-next-line no-undef
  var base64 = btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap))));
  var data = "sourceMappingURL=data:application/json;charset=utf-8;base64,".concat(base64);
  return "/*# ".concat(data, " */");
}

/***/ }),
/* 6 */
/***/ (function(module, exports, __webpack_require__) {

/*
	MIT License http://www.opensource.org/licenses/mit-license.php
	Author Tobias Koppers @sokra
*/

var stylesInDom = {};

var	memoize = function (fn) {
	var memo;

	return function () {
		if (typeof memo === "undefined") memo = fn.apply(this, arguments);
		return memo;
	};
};

var isOldIE = memoize(function () {
	// Test for IE <= 9 as proposed by Browserhacks
	// @see http://browserhacks.com/#hack-e71d8692f65334173fee715c222cb805
	// Tests for existence of standard globals is to allow style-loader
	// to operate correctly into non-standard environments
	// @see https://github.com/webpack-contrib/style-loader/issues/177
	return window && document && document.all && !window.atob;
});

var getElement = (function (fn) {
	var memo = {};

	return function(selector) {
		if (typeof memo[selector] === "undefined") {
			var styleTarget = fn.call(this, selector);
			// Special case to return head of iframe instead of iframe itself
			if (styleTarget instanceof window.HTMLIFrameElement) {
				try {
					// This will throw an exception if access to iframe is blocked
					// due to cross-origin restrictions
					styleTarget = styleTarget.contentDocument.head;
				} catch(e) {
					styleTarget = null;
				}
			}
			memo[selector] = styleTarget;
		}
		return memo[selector]
	};
})(function (target) {
	return document.querySelector(target)
});

var singleton = null;
var	singletonCounter = 0;
var	stylesInsertedAtTop = [];

var	fixUrls = __webpack_require__(7);

module.exports = function(list, options) {
	if (typeof DEBUG !== "undefined" && DEBUG) {
		if (typeof document !== "object") throw new Error("The style-loader cannot be used in a non-browser environment");
	}

	options = options || {};

	options.attrs = typeof options.attrs === "object" ? options.attrs : {};

	// Force single-tag solution on IE6-9, which has a hard limit on the # of <style>
	// tags it will allow on a page
	if (!options.singleton && typeof options.singleton !== "boolean") options.singleton = isOldIE();

	// By default, add <style> tags to the <head> element
	if (!options.insertInto) options.insertInto = "head";

	// By default, add <style> tags to the bottom of the target
	if (!options.insertAt) options.insertAt = "bottom";

	var styles = listToStyles(list, options);

	addStylesToDom(styles, options);

	return function update (newList) {
		var mayRemove = [];

		for (var i = 0; i < styles.length; i++) {
			var item = styles[i];
			var domStyle = stylesInDom[item.id];

			domStyle.refs--;
			mayRemove.push(domStyle);
		}

		if(newList) {
			var newStyles = listToStyles(newList, options);
			addStylesToDom(newStyles, options);
		}

		for (var i = 0; i < mayRemove.length; i++) {
			var domStyle = mayRemove[i];

			if(domStyle.refs === 0) {
				for (var j = 0; j < domStyle.parts.length; j++) domStyle.parts[j]();

				delete stylesInDom[domStyle.id];
			}
		}
	};
};

function addStylesToDom (styles, options) {
	for (var i = 0; i < styles.length; i++) {
		var item = styles[i];
		var domStyle = stylesInDom[item.id];

		if(domStyle) {
			domStyle.refs++;

			for(var j = 0; j < domStyle.parts.length; j++) {
				domStyle.parts[j](item.parts[j]);
			}

			for(; j < item.parts.length; j++) {
				domStyle.parts.push(addStyle(item.parts[j], options));
			}
		} else {
			var parts = [];

			for(var j = 0; j < item.parts.length; j++) {
				parts.push(addStyle(item.parts[j], options));
			}

			stylesInDom[item.id] = {id: item.id, refs: 1, parts: parts};
		}
	}
}

function listToStyles (list, options) {
	var styles = [];
	var newStyles = {};

	for (var i = 0; i < list.length; i++) {
		var item = list[i];
		var id = options.base ? item[0] + options.base : item[0];
		var css = item[1];
		var media = item[2];
		var sourceMap = item[3];
		var part = {css: css, media: media, sourceMap: sourceMap};

		if(!newStyles[id]) styles.push(newStyles[id] = {id: id, parts: [part]});
		else newStyles[id].parts.push(part);
	}

	return styles;
}

function insertStyleElement (options, style) {
	var target = getElement(options.insertInto)

	if (!target) {
		throw new Error("Couldn't find a style target. This probably means that the value for the 'insertInto' parameter is invalid.");
	}

	var lastStyleElementInsertedAtTop = stylesInsertedAtTop[stylesInsertedAtTop.length - 1];

	if (options.insertAt === "top") {
		if (!lastStyleElementInsertedAtTop) {
			target.insertBefore(style, target.firstChild);
		} else if (lastStyleElementInsertedAtTop.nextSibling) {
			target.insertBefore(style, lastStyleElementInsertedAtTop.nextSibling);
		} else {
			target.appendChild(style);
		}
		stylesInsertedAtTop.push(style);
	} else if (options.insertAt === "bottom") {
		target.appendChild(style);
	} else if (typeof options.insertAt === "object" && options.insertAt.before) {
		var nextSibling = getElement(options.insertInto + " " + options.insertAt.before);
		target.insertBefore(style, nextSibling);
	} else {
		throw new Error("[Style Loader]\n\n Invalid value for parameter 'insertAt' ('options.insertAt') found.\n Must be 'top', 'bottom', or Object.\n (https://github.com/webpack-contrib/style-loader#insertat)\n");
	}
}

function removeStyleElement (style) {
	if (style.parentNode === null) return false;
	style.parentNode.removeChild(style);

	var idx = stylesInsertedAtTop.indexOf(style);
	if(idx >= 0) {
		stylesInsertedAtTop.splice(idx, 1);
	}
}

function createStyleElement (options) {
	var style = document.createElement("style");

	options.attrs.type = "text/css";

	addAttrs(style, options.attrs);
	insertStyleElement(options, style);

	return style;
}

function createLinkElement (options) {
	var link = document.createElement("link");

	options.attrs.type = "text/css";
	options.attrs.rel = "stylesheet";

	addAttrs(link, options.attrs);
	insertStyleElement(options, link);

	return link;
}

function addAttrs (el, attrs) {
	Object.keys(attrs).forEach(function (key) {
		el.setAttribute(key, attrs[key]);
	});
}

function addStyle (obj, options) {
	var style, update, remove, result;

	// If a transform function was defined, run it on the css
	if (options.transform && obj.css) {
	    result = options.transform(obj.css);

	    if (result) {
	    	// If transform returns a value, use that instead of the original css.
	    	// This allows running runtime transformations on the css.
	    	obj.css = result;
	    } else {
	    	// If the transform function returns a falsy value, don't add this css.
	    	// This allows conditional loading of css
	    	return function() {
	    		// noop
	    	};
	    }
	}

	if (options.singleton) {
		var styleIndex = singletonCounter++;

		style = singleton || (singleton = createStyleElement(options));

		update = applyToSingletonTag.bind(null, style, styleIndex, false);
		remove = applyToSingletonTag.bind(null, style, styleIndex, true);

	} else if (
		obj.sourceMap &&
		typeof URL === "function" &&
		typeof URL.createObjectURL === "function" &&
		typeof URL.revokeObjectURL === "function" &&
		typeof Blob === "function" &&
		typeof btoa === "function"
	) {
		style = createLinkElement(options);
		update = updateLink.bind(null, style, options);
		remove = function () {
			removeStyleElement(style);

			if(style.href) URL.revokeObjectURL(style.href);
		};
	} else {
		style = createStyleElement(options);
		update = applyToTag.bind(null, style);
		remove = function () {
			removeStyleElement(style);
		};
	}

	update(obj);

	return function updateStyle (newObj) {
		if (newObj) {
			if (
				newObj.css === obj.css &&
				newObj.media === obj.media &&
				newObj.sourceMap === obj.sourceMap
			) {
				return;
			}

			update(obj = newObj);
		} else {
			remove();
		}
	};
}

var replaceText = (function () {
	var textStore = [];

	return function (index, replacement) {
		textStore[index] = replacement;

		return textStore.filter(Boolean).join('\n');
	};
})();

function applyToSingletonTag (style, index, remove, obj) {
	var css = remove ? "" : obj.css;

	if (style.styleSheet) {
		style.styleSheet.cssText = replaceText(index, css);
	} else {
		var cssNode = document.createTextNode(css);
		var childNodes = style.childNodes;

		if (childNodes[index]) style.removeChild(childNodes[index]);

		if (childNodes.length) {
			style.insertBefore(cssNode, childNodes[index]);
		} else {
			style.appendChild(cssNode);
		}
	}
}

function applyToTag (style, obj) {
	var css = obj.css;
	var media = obj.media;

	if(media) {
		style.setAttribute("media", media)
	}

	if(style.styleSheet) {
		style.styleSheet.cssText = css;
	} else {
		while(style.firstChild) {
			style.removeChild(style.firstChild);
		}

		style.appendChild(document.createTextNode(css));
	}
}

function updateLink (link, options, obj) {
	var css = obj.css;
	var sourceMap = obj.sourceMap;

	/*
		If convertToAbsoluteUrls isn't defined, but sourcemaps are enabled
		and there is no publicPath defined then lets turn convertToAbsoluteUrls
		on by default.  Otherwise default to the convertToAbsoluteUrls option
		directly
	*/
	var autoFixUrls = options.convertToAbsoluteUrls === undefined && sourceMap;

	if (options.convertToAbsoluteUrls || autoFixUrls) {
		css = fixUrls(css);
	}

	if (sourceMap) {
		// http://stackoverflow.com/a/26603875
		css += "\n/*# sourceMappingURL=data:application/json;base64," + btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap)))) + " */";
	}

	var blob = new Blob([css], { type: "text/css" });

	var oldSrc = link.href;

	link.href = URL.createObjectURL(blob);

	if(oldSrc) URL.revokeObjectURL(oldSrc);
}


/***/ }),
/* 7 */
/***/ (function(module, exports) {


/**
 * When source maps are enabled, `style-loader` uses a link element with a data-uri to
 * embed the css on the page. This breaks all relative urls because now they are relative to a
 * bundle instead of the current page.
 *
 * One solution is to only use full urls, but that may be impossible.
 *
 * Instead, this function "fixes" the relative urls to be absolute according to the current page location.
 *
 * A rudimentary test suite is located at `test/fixUrls.js` and can be run via the `npm test` command.
 *
 */

module.exports = function (css) {
  // get current location
  var location = typeof window !== "undefined" && window.location;

  if (!location) {
    throw new Error("fixUrls requires window.location");
  }

	// blank or null?
	if (!css || typeof css !== "string") {
	  return css;
  }

  var baseUrl = location.protocol + "//" + location.host;
  var currentDir = baseUrl + location.pathname.replace(/\/[^\/]*$/, "/");

	// convert each url(...)
	/*
	This regular expression is just a way to recursively match brackets within
	a string.

	 /url\s*\(  = Match on the word "url" with any whitespace after it and then a parens
	   (  = Start a capturing group
	     (?:  = Start a non-capturing group
	         [^)(]  = Match anything that isn't a parentheses
	         |  = OR
	         \(  = Match a start parentheses
	             (?:  = Start another non-capturing groups
	                 [^)(]+  = Match anything that isn't a parentheses
	                 |  = OR
	                 \(  = Match a start parentheses
	                     [^)(]*  = Match anything that isn't a parentheses
	                 \)  = Match a end parentheses
	             )  = End Group
              *\) = Match anything and then a close parens
          )  = Close non-capturing group
          *  = Match anything
       )  = Close capturing group
	 \)  = Match a close parens

	 /gi  = Get all matches, not the first.  Be case insensitive.
	 */
	var fixedCss = css.replace(/url\s*\(((?:[^)(]|\((?:[^)(]+|\([^)(]*\))*\))*)\)/gi, function(fullMatch, origUrl) {
		// strip quotes (if they exist)
		var unquotedOrigUrl = origUrl
			.trim()
			.replace(/^"(.*)"$/, function(o, $1){ return $1; })
			.replace(/^'(.*)'$/, function(o, $1){ return $1; });

		// already a full url? no change
		if (/^(#|data:|http:\/\/|https:\/\/|file:\/\/\/)/i.test(unquotedOrigUrl)) {
		  return fullMatch;
		}

		// convert the url to a full url
		var newUrl;

		if (unquotedOrigUrl.indexOf("//") === 0) {
		  	//TODO: should we add protocol?
			newUrl = unquotedOrigUrl;
		} else if (unquotedOrigUrl.indexOf("/") === 0) {
			// path should be relative to the base url
			newUrl = baseUrl + unquotedOrigUrl; // already starts with '/'
		} else {
			// path should be relative to current directory
			newUrl = currentDir + unquotedOrigUrl.replace(/^\.\//, ""); // Strip leading './'
		}

		// send back the fixed url(...)
		return "url(" + JSON.stringify(newUrl) + ")";
	});

	// send back the fixed css
	return fixedCss;
};


/***/ }),
/* 8 */
/***/ (function(module, exports, __webpack_require__) {

"use strict";
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.



function _regeneratorRuntime() { "use strict"; /*! regenerator-runtime -- Copyright (c) 2014-present, Facebook, Inc. -- license (MIT): https://github.com/facebook/regenerator/blob/main/LICENSE */ _regeneratorRuntime = function _regeneratorRuntime() { return e; }; var t, e = {}, r = Object.prototype, n = r.hasOwnProperty, o = Object.defineProperty || function (t, e, r) { t[e] = r.value; }, i = "function" == typeof Symbol ? Symbol : {}, a = i.iterator || "@@iterator", c = i.asyncIterator || "@@asyncIterator", u = i.toStringTag || "@@toStringTag"; function define(t, e, r) { return Object.defineProperty(t, e, { value: r, enumerable: !0, configurable: !0, writable: !0 }), t[e]; } try { define({}, ""); } catch (t) { define = function define(t, e, r) { return t[e] = r; }; } function wrap(t, e, r, n) { var i = e && e.prototype instanceof Generator ? e : Generator, a = Object.create(i.prototype), c = new Context(n || []); return o(a, "_invoke", { value: makeInvokeMethod(t, r, c) }), a; } function tryCatch(t, e, r) { try { return { type: "normal", arg: t.call(e, r) }; } catch (t) { return { type: "throw", arg: t }; } } e.wrap = wrap; var h = "suspendedStart", l = "suspendedYield", f = "executing", s = "completed", y = {}; function Generator() {} function GeneratorFunction() {} function GeneratorFunctionPrototype() {} var p = {}; define(p, a, function () { return this; }); var d = Object.getPrototypeOf, v = d && d(d(values([]))); v && v !== r && n.call(v, a) && (p = v); var g = GeneratorFunctionPrototype.prototype = Generator.prototype = Object.create(p); function defineIteratorMethods(t) { ["next", "throw", "return"].forEach(function (e) { define(t, e, function (t) { return this._invoke(e, t); }); }); } function AsyncIterator(t, e) { function invoke(r, o, i, a) { var c = tryCatch(t[r], t, o); if ("throw" !== c.type) { var u = c.arg, h = u.value; return h && "object" == _typeof(h) && n.call(h, "__await") ? e.resolve(h.__await).then(function (t) { invoke("next", t, i, a); }, function (t) { invoke("throw", t, i, a); }) : e.resolve(h).then(function (t) { u.value = t, i(u); }, function (t) { return invoke("throw", t, i, a); }); } a(c.arg); } var r; o(this, "_invoke", { value: function value(t, n) { function callInvokeWithMethodAndArg() { return new e(function (e, r) { invoke(t, n, e, r); }); } return r = r ? r.then(callInvokeWithMethodAndArg, callInvokeWithMethodAndArg) : callInvokeWithMethodAndArg(); } }); } function makeInvokeMethod(e, r, n) { var o = h; return function (i, a) { if (o === f) throw new Error("Generator is already running"); if (o === s) { if ("throw" === i) throw a; return { value: t, done: !0 }; } for (n.method = i, n.arg = a;;) { var c = n.delegate; if (c) { var u = maybeInvokeDelegate(c, n); if (u) { if (u === y) continue; return u; } } if ("next" === n.method) n.sent = n._sent = n.arg;else if ("throw" === n.method) { if (o === h) throw o = s, n.arg; n.dispatchException(n.arg); } else "return" === n.method && n.abrupt("return", n.arg); o = f; var p = tryCatch(e, r, n); if ("normal" === p.type) { if (o = n.done ? s : l, p.arg === y) continue; return { value: p.arg, done: n.done }; } "throw" === p.type && (o = s, n.method = "throw", n.arg = p.arg); } }; } function maybeInvokeDelegate(e, r) { var n = r.method, o = e.iterator[n]; if (o === t) return r.delegate = null, "throw" === n && e.iterator["return"] && (r.method = "return", r.arg = t, maybeInvokeDelegate(e, r), "throw" === r.method) || "return" !== n && (r.method = "throw", r.arg = new TypeError("The iterator does not provide a '" + n + "' method")), y; var i = tryCatch(o, e.iterator, r.arg); if ("throw" === i.type) return r.method = "throw", r.arg = i.arg, r.delegate = null, y; var a = i.arg; return a ? a.done ? (r[e.resultName] = a.value, r.next = e.nextLoc, "return" !== r.method && (r.method = "next", r.arg = t), r.delegate = null, y) : a : (r.method = "throw", r.arg = new TypeError("iterator result is not an object"), r.delegate = null, y); } function pushTryEntry(t) { var e = { tryLoc: t[0] }; 1 in t && (e.catchLoc = t[1]), 2 in t && (e.finallyLoc = t[2], e.afterLoc = t[3]), this.tryEntries.push(e); } function resetTryEntry(t) { var e = t.completion || {}; e.type = "normal", delete e.arg, t.completion = e; } function Context(t) { this.tryEntries = [{ tryLoc: "root" }], t.forEach(pushTryEntry, this), this.reset(!0); } function values(e) { if (e || "" === e) { var r = e[a]; if (r) return r.call(e); if ("function" == typeof e.next) return e; if (!isNaN(e.length)) { var o = -1, i = function next() { for (; ++o < e.length;) if (n.call(e, o)) return next.value = e[o], next.done = !1, next; return next.value = t, next.done = !0, next; }; return i.next = i; } } throw new TypeError(_typeof(e) + " is not iterable"); } return GeneratorFunction.prototype = GeneratorFunctionPrototype, o(g, "constructor", { value: GeneratorFunctionPrototype, configurable: !0 }), o(GeneratorFunctionPrototype, "constructor", { value: GeneratorFunction, configurable: !0 }), GeneratorFunction.displayName = define(GeneratorFunctionPrototype, u, "GeneratorFunction"), e.isGeneratorFunction = function (t) { var e = "function" == typeof t && t.constructor; return !!e && (e === GeneratorFunction || "GeneratorFunction" === (e.displayName || e.name)); }, e.mark = function (t) { return Object.setPrototypeOf ? Object.setPrototypeOf(t, GeneratorFunctionPrototype) : (t.__proto__ = GeneratorFunctionPrototype, define(t, u, "GeneratorFunction")), t.prototype = Object.create(g), t; }, e.awrap = function (t) { return { __await: t }; }, defineIteratorMethods(AsyncIterator.prototype), define(AsyncIterator.prototype, c, function () { return this; }), e.AsyncIterator = AsyncIterator, e.async = function (t, r, n, o, i) { void 0 === i && (i = Promise); var a = new AsyncIterator(wrap(t, r, n, o), i); return e.isGeneratorFunction(r) ? a : a.next().then(function (t) { return t.done ? t.value : a.next(); }); }, defineIteratorMethods(g), define(g, u, "Generator"), define(g, a, function () { return this; }), define(g, "toString", function () { return "[object Generator]"; }), e.keys = function (t) { var e = Object(t), r = []; for (var n in e) r.push(n); return r.reverse(), function next() { for (; r.length;) { var t = r.pop(); if (t in e) return next.value = t, next.done = !1, next; } return next.done = !0, next; }; }, e.values = values, Context.prototype = { constructor: Context, reset: function reset(e) { if (this.prev = 0, this.next = 0, this.sent = this._sent = t, this.done = !1, this.delegate = null, this.method = "next", this.arg = t, this.tryEntries.forEach(resetTryEntry), !e) for (var r in this) "t" === r.charAt(0) && n.call(this, r) && !isNaN(+r.slice(1)) && (this[r] = t); }, stop: function stop() { this.done = !0; var t = this.tryEntries[0].completion; if ("throw" === t.type) throw t.arg; return this.rval; }, dispatchException: function dispatchException(e) { if (this.done) throw e; var r = this; function handle(n, o) { return a.type = "throw", a.arg = e, r.next = n, o && (r.method = "next", r.arg = t), !!o; } for (var o = this.tryEntries.length - 1; o >= 0; --o) { var i = this.tryEntries[o], a = i.completion; if ("root" === i.tryLoc) return handle("end"); if (i.tryLoc <= this.prev) { var c = n.call(i, "catchLoc"), u = n.call(i, "finallyLoc"); if (c && u) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } else if (c) { if (this.prev < i.catchLoc) return handle(i.catchLoc, !0); } else { if (!u) throw new Error("try statement without catch or finally"); if (this.prev < i.finallyLoc) return handle(i.finallyLoc); } } } }, abrupt: function abrupt(t, e) { for (var r = this.tryEntries.length - 1; r >= 0; --r) { var o = this.tryEntries[r]; if (o.tryLoc <= this.prev && n.call(o, "finallyLoc") && this.prev < o.finallyLoc) { var i = o; break; } } i && ("break" === t || "continue" === t) && i.tryLoc <= e && e <= i.finallyLoc && (i = null); var a = i ? i.completion : {}; return a.type = t, a.arg = e, i ? (this.method = "next", this.next = i.finallyLoc, y) : this.complete(a); }, complete: function complete(t, e) { if ("throw" === t.type) throw t.arg; return "break" === t.type || "continue" === t.type ? this.next = t.arg : "return" === t.type ? (this.rval = this.arg = t.arg, this.method = "return", this.next = "end") : "normal" === t.type && e && (this.next = e), y; }, finish: function finish(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.finallyLoc === t) return this.complete(r.completion, r.afterLoc), resetTryEntry(r), y; } }, "catch": function _catch(t) { for (var e = this.tryEntries.length - 1; e >= 0; --e) { var r = this.tryEntries[e]; if (r.tryLoc === t) { var n = r.completion; if ("throw" === n.type) { var o = n.arg; resetTryEntry(r); } return o; } } throw new Error("illegal catch attempt"); }, delegateYield: function delegateYield(e, r, n) { return this.delegate = { iterator: values(e), resultName: r, nextLoc: n }, "next" === this.method && (this.arg = t), y; } }, e; }
function asyncGeneratorStep(gen, resolve, reject, _next, _throw, key, arg) { try { var info = gen[key](arg); var value = info.value; } catch (error) { reject(error); return; } if (info.done) { resolve(value); } else { Promise.resolve(value).then(_next, _throw); } }
function _asyncToGenerator(fn) { return function () { var self = this, args = arguments; return new Promise(function (resolve, reject) { var gen = fn.apply(self, args); function _next(value) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "next", value); } function _throw(err) { asyncGeneratorStep(gen, resolve, reject, _next, _throw, "throw", err); } _next(undefined); }); }; }
function _typeof(o) { "@babel/helpers - typeof"; return _typeof = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function (o) { return typeof o; } : function (o) { return o && "function" == typeof Symbol && o.constructor === Symbol && o !== Symbol.prototype ? "symbol" : typeof o; }, _typeof(o); }
function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, _toPropertyKey(descriptor.key), descriptor); } }
function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }
function _toPropertyKey(arg) { var key = _toPrimitive(arg, "string"); return _typeof(key) === "symbol" ? key : String(key); }
function _toPrimitive(input, hint) { if (_typeof(input) !== "object" || input === null) return input; var prim = input[Symbol.toPrimitive]; if (prim !== undefined) { var res = prim.call(input, hint || "default"); if (_typeof(res) !== "object") return res; throw new TypeError("@@toPrimitive must return a primitive value."); } return (hint === "string" ? String : Number)(input); }
(function () {
  var enosys = function enosys() {
    var err = new Error("not implemented");
    err.code = "ENOSYS";
    return err;
  };
  if (!globalThis.fs) {
    var outputBuf = "";
    globalThis.fs = {
      constants: {
        O_WRONLY: -1,
        O_RDWR: -1,
        O_CREAT: -1,
        O_TRUNC: -1,
        O_APPEND: -1,
        O_EXCL: -1
      },
      // unused
      writeSync: function writeSync(fd, buf) {
        outputBuf += decoder.decode(buf);
        var nl = outputBuf.lastIndexOf("\n");
        if (nl != -1) {
          console.log(outputBuf.substring(0, nl));
          outputBuf = outputBuf.substring(nl + 1);
        }
        return buf.length;
      },
      write: function write(fd, buf, offset, length, position, callback) {
        if (offset !== 0 || length !== buf.length || position !== null) {
          callback(enosys());
          return;
        }
        var n = this.writeSync(fd, buf);
        callback(null, n);
      },
      chmod: function chmod(path, mode, callback) {
        callback(enosys());
      },
      chown: function chown(path, uid, gid, callback) {
        callback(enosys());
      },
      close: function close(fd, callback) {
        callback(enosys());
      },
      fchmod: function fchmod(fd, mode, callback) {
        callback(enosys());
      },
      fchown: function fchown(fd, uid, gid, callback) {
        callback(enosys());
      },
      fstat: function fstat(fd, callback) {
        callback(enosys());
      },
      fsync: function fsync(fd, callback) {
        callback(null);
      },
      ftruncate: function ftruncate(fd, length, callback) {
        callback(enosys());
      },
      lchown: function lchown(path, uid, gid, callback) {
        callback(enosys());
      },
      link: function link(path, _link, callback) {
        callback(enosys());
      },
      lstat: function lstat(path, callback) {
        callback(enosys());
      },
      mkdir: function mkdir(path, perm, callback) {
        callback(enosys());
      },
      open: function open(path, flags, mode, callback) {
        callback(enosys());
      },
      read: function read(fd, buffer, offset, length, position, callback) {
        callback(enosys());
      },
      readdir: function readdir(path, callback) {
        callback(enosys());
      },
      readlink: function readlink(path, callback) {
        callback(enosys());
      },
      rename: function rename(from, to, callback) {
        callback(enosys());
      },
      rmdir: function rmdir(path, callback) {
        callback(enosys());
      },
      stat: function stat(path, callback) {
        callback(enosys());
      },
      symlink: function symlink(path, link, callback) {
        callback(enosys());
      },
      truncate: function truncate(path, length, callback) {
        callback(enosys());
      },
      unlink: function unlink(path, callback) {
        callback(enosys());
      },
      utimes: function utimes(path, atime, mtime, callback) {
        callback(enosys());
      }
    };
  }
  if (!globalThis.process) {
    globalThis.process = {
      getuid: function getuid() {
        return -1;
      },
      getgid: function getgid() {
        return -1;
      },
      geteuid: function geteuid() {
        return -1;
      },
      getegid: function getegid() {
        return -1;
      },
      getgroups: function getgroups() {
        throw enosys();
      },
      pid: -1,
      ppid: -1,
      umask: function umask() {
        throw enosys();
      },
      cwd: function cwd() {
        throw enosys();
      },
      chdir: function chdir() {
        throw enosys();
      }
    };
  }
  if (!globalThis.crypto) {
    throw new Error("globalThis.crypto is not available, polyfill required (crypto.getRandomValues only)");
  }
  if (!globalThis.performance) {
    throw new Error("globalThis.performance is not available, polyfill required (performance.now only)");
  }
  if (!globalThis.TextEncoder) {
    throw new Error("globalThis.TextEncoder is not available, polyfill required");
  }
  if (!globalThis.TextDecoder) {
    throw new Error("globalThis.TextDecoder is not available, polyfill required");
  }
  var encoder = new TextEncoder("utf-8");
  var decoder = new TextDecoder("utf-8");
  globalThis.Go = /*#__PURE__*/function () {
    function _class() {
      var _this = this;
      _classCallCheck(this, _class);
      this.argv = ["js"];
      this.env = {};
      this.exit = function (code) {
        if (code !== 0) {
          console.warn("exit code:", code);
        }
      };
      this._exitPromise = new Promise(function (resolve) {
        _this._resolveExitPromise = resolve;
      });
      this._pendingEvent = null;
      this._scheduledTimeouts = new Map();
      this._nextCallbackTimeoutID = 1;
      var setInt64 = function setInt64(addr, v) {
        _this.mem.setUint32(addr + 0, v, true);
        _this.mem.setUint32(addr + 4, Math.floor(v / 4294967296), true);
      };
      var setInt32 = function setInt32(addr, v) {
        _this.mem.setUint32(addr + 0, v, true);
      };
      var getInt64 = function getInt64(addr) {
        var low = _this.mem.getUint32(addr + 0, true);
        var high = _this.mem.getInt32(addr + 4, true);
        return low + high * 4294967296;
      };
      var loadValue = function loadValue(addr) {
        var f = _this.mem.getFloat64(addr, true);
        if (f === 0) {
          return undefined;
        }
        if (!isNaN(f)) {
          return f;
        }
        var id = _this.mem.getUint32(addr, true);
        return _this._values[id];
      };
      var storeValue = function storeValue(addr, v) {
        var nanHead = 0x7FF80000;
        if (typeof v === "number" && v !== 0) {
          if (isNaN(v)) {
            _this.mem.setUint32(addr + 4, nanHead, true);
            _this.mem.setUint32(addr, 0, true);
            return;
          }
          _this.mem.setFloat64(addr, v, true);
          return;
        }
        if (v === undefined) {
          _this.mem.setFloat64(addr, 0, true);
          return;
        }
        var id = _this._ids.get(v);
        if (id === undefined) {
          id = _this._idPool.pop();
          if (id === undefined) {
            id = _this._values.length;
          }
          _this._values[id] = v;
          _this._goRefCounts[id] = 0;
          _this._ids.set(v, id);
        }
        _this._goRefCounts[id]++;
        var typeFlag = 0;
        switch (_typeof(v)) {
          case "object":
            if (v !== null) {
              typeFlag = 1;
            }
            break;
          case "string":
            typeFlag = 2;
            break;
          case "symbol":
            typeFlag = 3;
            break;
          case "function":
            typeFlag = 4;
            break;
        }
        _this.mem.setUint32(addr + 4, nanHead | typeFlag, true);
        _this.mem.setUint32(addr, id, true);
      };
      var loadSlice = function loadSlice(addr) {
        var array = getInt64(addr + 0);
        var len = getInt64(addr + 8);
        return new Uint8Array(_this._inst.exports.mem.buffer, array, len);
      };
      var loadSliceOfValues = function loadSliceOfValues(addr) {
        var array = getInt64(addr + 0);
        var len = getInt64(addr + 8);
        var a = new Array(len);
        for (var i = 0; i < len; i++) {
          a[i] = loadValue(array + i * 8);
        }
        return a;
      };
      var loadString = function loadString(addr) {
        var saddr = getInt64(addr + 0);
        var len = getInt64(addr + 8);
        return decoder.decode(new DataView(_this._inst.exports.mem.buffer, saddr, len));
      };
      var timeOrigin = Date.now() - performance.now();
      this.importObject = {
        _gotest: {
          add: function add(a, b) {
            return a + b;
          }
        },
        gojs: {
          // Go's SP does not change as long as no Go code is running. Some operations (e.g. calls, getters and setters)
          // may synchronously trigger a Go event handler. This makes Go code get executed in the middle of the imported
          // function. A goroutine can switch to a new stack if the current stack is too small (see morestack function).
          // This changes the SP, thus we have to update the SP used by the imported function.

          // func wasmExit(code int32)
          "runtime.wasmExit": function runtimeWasmExit(sp) {
            sp >>>= 0;
            var code = _this.mem.getInt32(sp + 8, true);
            _this.exited = true;
            delete _this._inst;
            delete _this._values;
            delete _this._goRefCounts;
            delete _this._ids;
            delete _this._idPool;
            _this.exit(code);
          },
          // func wasmWrite(fd uintptr, p unsafe.Pointer, n int32)
          "runtime.wasmWrite": function runtimeWasmWrite(sp) {
            sp >>>= 0;
            var fd = getInt64(sp + 8);
            var p = getInt64(sp + 16);
            var n = _this.mem.getInt32(sp + 24, true);
            fs.writeSync(fd, new Uint8Array(_this._inst.exports.mem.buffer, p, n));
          },
          // func resetMemoryDataView()
          "runtime.resetMemoryDataView": function runtimeResetMemoryDataView(sp) {
            sp >>>= 0;
            _this.mem = new DataView(_this._inst.exports.mem.buffer);
          },
          // func nanotime1() int64
          "runtime.nanotime1": function runtimeNanotime1(sp) {
            sp >>>= 0;
            setInt64(sp + 8, (timeOrigin + performance.now()) * 1000000);
          },
          // func walltime() (sec int64, nsec int32)
          "runtime.walltime": function runtimeWalltime(sp) {
            sp >>>= 0;
            var msec = new Date().getTime();
            setInt64(sp + 8, msec / 1000);
            _this.mem.setInt32(sp + 16, msec % 1000 * 1000000, true);
          },
          // func scheduleTimeoutEvent(delay int64) int32
          "runtime.scheduleTimeoutEvent": function runtimeScheduleTimeoutEvent(sp) {
            sp >>>= 0;
            var id = _this._nextCallbackTimeoutID;
            _this._nextCallbackTimeoutID++;
            _this._scheduledTimeouts.set(id, setTimeout(function () {
              _this._resume();
              while (_this._scheduledTimeouts.has(id)) {
                // for some reason Go failed to register the timeout event, log and try again
                // (temporary workaround for https://github.com/golang/go/issues/28975)
                console.warn("scheduleTimeoutEvent: missed timeout event");
                _this._resume();
              }
            }, getInt64(sp + 8)));
            _this.mem.setInt32(sp + 16, id, true);
          },
          // func clearTimeoutEvent(id int32)
          "runtime.clearTimeoutEvent": function runtimeClearTimeoutEvent(sp) {
            sp >>>= 0;
            var id = _this.mem.getInt32(sp + 8, true);
            clearTimeout(_this._scheduledTimeouts.get(id));
            _this._scheduledTimeouts["delete"](id);
          },
          // func getRandomData(r []byte)
          "runtime.getRandomData": function runtimeGetRandomData(sp) {
            sp >>>= 0;
            crypto.getRandomValues(loadSlice(sp + 8));
          },
          // func finalizeRef(v ref)
          "syscall/js.finalizeRef": function syscallJsFinalizeRef(sp) {
            sp >>>= 0;
            var id = _this.mem.getUint32(sp + 8, true);
            _this._goRefCounts[id]--;
            if (_this._goRefCounts[id] === 0) {
              var v = _this._values[id];
              _this._values[id] = null;
              _this._ids["delete"](v);
              _this._idPool.push(id);
            }
          },
          // func stringVal(value string) ref
          "syscall/js.stringVal": function syscallJsStringVal(sp) {
            sp >>>= 0;
            storeValue(sp + 24, loadString(sp + 8));
          },
          // func valueGet(v ref, p string) ref
          "syscall/js.valueGet": function syscallJsValueGet(sp) {
            sp >>>= 0;
            var result = Reflect.get(loadValue(sp + 8), loadString(sp + 16));
            sp = _this._inst.exports.getsp() >>> 0; // see comment above
            storeValue(sp + 32, result);
          },
          // func valueSet(v ref, p string, x ref)
          "syscall/js.valueSet": function syscallJsValueSet(sp) {
            sp >>>= 0;
            Reflect.set(loadValue(sp + 8), loadString(sp + 16), loadValue(sp + 32));
          },
          // func valueDelete(v ref, p string)
          "syscall/js.valueDelete": function syscallJsValueDelete(sp) {
            sp >>>= 0;
            Reflect.deleteProperty(loadValue(sp + 8), loadString(sp + 16));
          },
          // func valueIndex(v ref, i int) ref
          "syscall/js.valueIndex": function syscallJsValueIndex(sp) {
            sp >>>= 0;
            storeValue(sp + 24, Reflect.get(loadValue(sp + 8), getInt64(sp + 16)));
          },
          // valueSetIndex(v ref, i int, x ref)
          "syscall/js.valueSetIndex": function syscallJsValueSetIndex(sp) {
            sp >>>= 0;
            Reflect.set(loadValue(sp + 8), getInt64(sp + 16), loadValue(sp + 24));
          },
          // func valueCall(v ref, m string, args []ref) (ref, bool)
          "syscall/js.valueCall": function syscallJsValueCall(sp) {
            sp >>>= 0;
            try {
              var v = loadValue(sp + 8);
              var m = Reflect.get(v, loadString(sp + 16));
              var args = loadSliceOfValues(sp + 32);
              var result = Reflect.apply(m, v, args);
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 56, result);
              _this.mem.setUint8(sp + 64, 1);
            } catch (err) {
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 56, err);
              _this.mem.setUint8(sp + 64, 0);
            }
          },
          // func valueInvoke(v ref, args []ref) (ref, bool)
          "syscall/js.valueInvoke": function syscallJsValueInvoke(sp) {
            sp >>>= 0;
            try {
              var v = loadValue(sp + 8);
              var args = loadSliceOfValues(sp + 16);
              var result = Reflect.apply(v, undefined, args);
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 40, result);
              _this.mem.setUint8(sp + 48, 1);
            } catch (err) {
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 40, err);
              _this.mem.setUint8(sp + 48, 0);
            }
          },
          // func valueNew(v ref, args []ref) (ref, bool)
          "syscall/js.valueNew": function syscallJsValueNew(sp) {
            sp >>>= 0;
            try {
              var v = loadValue(sp + 8);
              var args = loadSliceOfValues(sp + 16);
              var result = Reflect.construct(v, args);
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 40, result);
              _this.mem.setUint8(sp + 48, 1);
            } catch (err) {
              sp = _this._inst.exports.getsp() >>> 0; // see comment above
              storeValue(sp + 40, err);
              _this.mem.setUint8(sp + 48, 0);
            }
          },
          // func valueLength(v ref) int
          "syscall/js.valueLength": function syscallJsValueLength(sp) {
            sp >>>= 0;
            setInt64(sp + 16, parseInt(loadValue(sp + 8).length));
          },
          // valuePrepareString(v ref) (ref, int)
          "syscall/js.valuePrepareString": function syscallJsValuePrepareString(sp) {
            sp >>>= 0;
            var str = encoder.encode(String(loadValue(sp + 8)));
            storeValue(sp + 16, str);
            setInt64(sp + 24, str.length);
          },
          // valueLoadString(v ref, b []byte)
          "syscall/js.valueLoadString": function syscallJsValueLoadString(sp) {
            sp >>>= 0;
            var str = loadValue(sp + 8);
            loadSlice(sp + 16).set(str);
          },
          // func valueInstanceOf(v ref, t ref) bool
          "syscall/js.valueInstanceOf": function syscallJsValueInstanceOf(sp) {
            sp >>>= 0;
            _this.mem.setUint8(sp + 24, loadValue(sp + 8) instanceof loadValue(sp + 16) ? 1 : 0);
          },
          // func copyBytesToGo(dst []byte, src ref) (int, bool)
          "syscall/js.copyBytesToGo": function syscallJsCopyBytesToGo(sp) {
            sp >>>= 0;
            var dst = loadSlice(sp + 8);
            var src = loadValue(sp + 32);
            if (!(src instanceof Uint8Array || src instanceof Uint8ClampedArray)) {
              _this.mem.setUint8(sp + 48, 0);
              return;
            }
            var toCopy = src.subarray(0, dst.length);
            dst.set(toCopy);
            setInt64(sp + 40, toCopy.length);
            _this.mem.setUint8(sp + 48, 1);
          },
          // func copyBytesToJS(dst ref, src []byte) (int, bool)
          "syscall/js.copyBytesToJS": function syscallJsCopyBytesToJS(sp) {
            sp >>>= 0;
            var dst = loadValue(sp + 8);
            var src = loadSlice(sp + 16);
            if (!(dst instanceof Uint8Array || dst instanceof Uint8ClampedArray)) {
              _this.mem.setUint8(sp + 48, 0);
              return;
            }
            var toCopy = src.subarray(0, dst.length);
            dst.set(toCopy);
            setInt64(sp + 40, toCopy.length);
            _this.mem.setUint8(sp + 48, 1);
          },
          "debug": function debug(value) {
            console.log(value);
          }
        }
      };
    }
    _createClass(_class, [{
      key: "run",
      value: function () {
        var _run = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee(instance) {
          var _this2 = this;
          var offset, strPtr, argc, argvPtrs, keys, argv, wasmMinDataAddr;
          return _regeneratorRuntime().wrap(function _callee$(_context) {
            while (1) switch (_context.prev = _context.next) {
              case 0:
                if (instance instanceof WebAssembly.Instance) {
                  _context.next = 2;
                  break;
                }
                throw new Error("Go.run: WebAssembly.Instance expected");
              case 2:
                this._inst = instance;
                this.mem = new DataView(this._inst.exports.mem.buffer);
                this._values = [
                // JS values that Go currently has references to, indexed by reference id
                NaN, 0, null, true, false, globalThis, this];
                this._goRefCounts = new Array(this._values.length).fill(Infinity); // number of references that Go has to a JS value, indexed by reference id
                this._ids = new Map([
                // mapping from JS values to reference ids
                [0, 1], [null, 2], [true, 3], [false, 4], [globalThis, 5], [this, 6]]);
                this._idPool = []; // unused ids that have been garbage collected
                this.exited = false; // whether the Go program has exited

                // Pass command line arguments and environment variables to WebAssembly by writing them to the linear memory.
                offset = 4096;
                strPtr = function strPtr(str) {
                  var ptr = offset;
                  var bytes = encoder.encode(str + "\0");
                  new Uint8Array(_this2.mem.buffer, offset, bytes.length).set(bytes);
                  offset += bytes.length;
                  if (offset % 8 !== 0) {
                    offset += 8 - offset % 8;
                  }
                  return ptr;
                };
                argc = this.argv.length;
                argvPtrs = [];
                this.argv.forEach(function (arg) {
                  argvPtrs.push(strPtr(arg));
                });
                argvPtrs.push(0);
                keys = Object.keys(this.env).sort();
                keys.forEach(function (key) {
                  argvPtrs.push(strPtr("".concat(key, "=").concat(_this2.env[key])));
                });
                argvPtrs.push(0);
                argv = offset;
                argvPtrs.forEach(function (ptr) {
                  _this2.mem.setUint32(offset, ptr, true);
                  _this2.mem.setUint32(offset + 4, 0, true);
                  offset += 8;
                });

                // The linker guarantees global data starts from at least wasmMinDataAddr.
                // Keep in sync with cmd/link/internal/ld/data.go:wasmMinDataAddr.
                wasmMinDataAddr = 4096 + 8192;
                if (!(offset >= wasmMinDataAddr)) {
                  _context.next = 23;
                  break;
                }
                throw new Error("total length of command line and environment variables exceeds limit");
              case 23:
                this._inst.exports.run(argc, argv);
                if (this.exited) {
                  this._resolveExitPromise();
                }
                _context.next = 27;
                return this._exitPromise;
              case 27:
              case "end":
                return _context.stop();
            }
          }, _callee, this);
        }));
        function run(_x) {
          return _run.apply(this, arguments);
        }
        return run;
      }()
    }, {
      key: "_resume",
      value: function _resume() {
        if (this.exited) {
          throw new Error("Go program has already exited");
        }
        this._inst.exports.resume();
        if (this.exited) {
          this._resolveExitPromise();
        }
      }
    }, {
      key: "_makeFuncWrapper",
      value: function _makeFuncWrapper(id) {
        var go = this;
        return function () {
          var event = {
            id: id,
            "this": this,
            args: arguments
          };
          go._pendingEvent = event;
          go._resume();
          return event.result;
        };
      }
    }]);
    return _class;
  }();
})();

/***/ })
/******/ ]);
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAgMzNjNmY2MWRlMTNjYjcxMDczNzQiLCJ3ZWJwYWNrOi8vLy4vc3JjL21haW4uanMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3NlcnZpY2VzLmpzIiwid2VicGFjazovLy8uL3NyYy92aWV3cy9sYi5qcyIsIndlYnBhY2s6Ly8vLi9zcmMvdmlld3MvbGIuY3NzPzYzYTciLCJ3ZWJwYWNrOi8vLy4vc3JjL3ZpZXdzL2xiLmNzcyIsIndlYnBhY2s6Ly8vLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L3J1bnRpbWUvYXBpLmpzIiwid2VicGFjazovLy8uL25vZGVfbW9kdWxlcy9zdHlsZS1sb2FkZXIvbGliL2FkZFN0eWxlcy5qcyIsIndlYnBhY2s6Ly8vLi9ub2RlX21vZHVsZXMvc3R5bGUtbG9hZGVyL2xpYi91cmxzLmpzIiwid2VicGFjazovLy8uL3NyYy92aWV3cy93YXNtX2V4ZWMuanMiXSwibmFtZXMiOlsiX3JlZ2VuZXJhdG9yUnVudGltZSIsImUiLCJ0IiwiciIsIk9iamVjdCIsInByb3RvdHlwZSIsIm4iLCJoYXNPd25Qcm9wZXJ0eSIsIm8iLCJkZWZpbmVQcm9wZXJ0eSIsInZhbHVlIiwiaSIsIlN5bWJvbCIsImEiLCJpdGVyYXRvciIsImMiLCJhc3luY0l0ZXJhdG9yIiwidSIsInRvU3RyaW5nVGFnIiwiZGVmaW5lIiwiZW51bWVyYWJsZSIsImNvbmZpZ3VyYWJsZSIsIndyaXRhYmxlIiwid3JhcCIsIkdlbmVyYXRvciIsImNyZWF0ZSIsIkNvbnRleHQiLCJtYWtlSW52b2tlTWV0aG9kIiwidHJ5Q2F0Y2giLCJ0eXBlIiwiYXJnIiwiY2FsbCIsImgiLCJsIiwiZiIsInMiLCJ5IiwiR2VuZXJhdG9yRnVuY3Rpb24iLCJHZW5lcmF0b3JGdW5jdGlvblByb3RvdHlwZSIsInAiLCJkIiwiZ2V0UHJvdG90eXBlT2YiLCJ2IiwidmFsdWVzIiwiZyIsImRlZmluZUl0ZXJhdG9yTWV0aG9kcyIsImZvckVhY2giLCJfaW52b2tlIiwiQXN5bmNJdGVyYXRvciIsImludm9rZSIsIl90eXBlb2YiLCJyZXNvbHZlIiwiX19hd2FpdCIsInRoZW4iLCJjYWxsSW52b2tlV2l0aE1ldGhvZEFuZEFyZyIsIkVycm9yIiwiZG9uZSIsIm1ldGhvZCIsImRlbGVnYXRlIiwibWF5YmVJbnZva2VEZWxlZ2F0ZSIsInNlbnQiLCJfc2VudCIsImRpc3BhdGNoRXhjZXB0aW9uIiwiYWJydXB0IiwiVHlwZUVycm9yIiwicmVzdWx0TmFtZSIsIm5leHQiLCJuZXh0TG9jIiwicHVzaFRyeUVudHJ5IiwidHJ5TG9jIiwiY2F0Y2hMb2MiLCJmaW5hbGx5TG9jIiwiYWZ0ZXJMb2MiLCJ0cnlFbnRyaWVzIiwicHVzaCIsInJlc2V0VHJ5RW50cnkiLCJjb21wbGV0aW9uIiwicmVzZXQiLCJpc05hTiIsImxlbmd0aCIsImRpc3BsYXlOYW1lIiwiaXNHZW5lcmF0b3JGdW5jdGlvbiIsImNvbnN0cnVjdG9yIiwibmFtZSIsIm1hcmsiLCJzZXRQcm90b3R5cGVPZiIsIl9fcHJvdG9fXyIsImF3cmFwIiwiYXN5bmMiLCJQcm9taXNlIiwia2V5cyIsInJldmVyc2UiLCJwb3AiLCJwcmV2IiwiY2hhckF0Iiwic2xpY2UiLCJzdG9wIiwicnZhbCIsImhhbmRsZSIsImNvbXBsZXRlIiwiZmluaXNoIiwiX2NhdGNoIiwiZGVsZWdhdGVZaWVsZCIsImFzeW5jR2VuZXJhdG9yU3RlcCIsImdlbiIsInJlamVjdCIsIl9uZXh0IiwiX3Rocm93Iiwia2V5IiwiaW5mbyIsImVycm9yIiwiX2FzeW5jVG9HZW5lcmF0b3IiLCJmbiIsInNlbGYiLCJhcmdzIiwiYXJndW1lbnRzIiwiYXBwbHkiLCJlcnIiLCJ1bmRlZmluZWQiLCJzdXBwb3J0ZWRBUEkiLCJnbyIsImFwcCIsIndpbmRvdyIsImNvbnNvbGUiLCJsb2ciLCJjb25maWd1cmF0aW9ucyIsInNvbWVEZWZhdWx0Q29uZmlndXJhdGlvbiIsImdsb2JhbE9iamVjdCIsInF1ZXVlIiwicSIsInRvTG93ZXJDYXNlIiwiZXh0ZW5kT2JqZWN0IiwiYXBpSGFuZGxlciIsIl94IiwiX3gyIiwiX2FwaUhhbmRsZXIiLCJfY2FsbGVlIiwiYXBpIiwicGFyYW1zIiwiX2NhbGxlZSQiLCJfY29udGV4dCIsImluZGV4T2YiLCJjb25jYXQiLCJ0MCIsImxvYWRXYXNtIiwiYWRkQnV0dG9uVG9UZXh0YXJlYSIsIndhcm4iLCJiIiwiX2xvYWRXYXNtIiwiX2NhbGxlZTIiLCJyZXNwb25zZSIsImJ1ZmZlciIsIl9nbyIsInJlc3VsdCIsIl9jYWxsZWUyJCIsIl9jb250ZXh0MiIsImZldGNoIiwiYXJyYXlCdWZmZXIiLCJHbyIsIldlYkFzc2VtYmx5IiwiaW5zdGFudGlhdGUiLCJpbXBvcnRPYmplY3QiLCJydW4iLCJpbnN0YW5jZSIsInBpbmciLCJjbGFzc05hbWUiLCJkaXZFbGVtZW50cyIsImRvY3VtZW50IiwicXVlcnlTZWxlY3RvckFsbCIsImRpdkVsZW1lbnQiLCJidXR0b25JbWFnZSIsImNyZWF0ZUVsZW1lbnQiLCJzcmMiLCJzdHlsZSIsIndpZHRoIiwiaW5kZXgiLCJNYXRoIiwicmFuZG9tIiwidG9TdHJpbmciLCJzdWJzdHIiLCJidXR0b25FbGVtZW50IiwiYXBwZW5kQ2hpbGQiLCJhZGRFdmVudExpc3RlbmVyIiwiYWRkVGV4dFRvRGl2IiwiYWxsIiwicG9zaXRpb24iLCJib3R0b20iLCJyaWdodCIsInRyYW5zaXRpb24iLCJjdXJzb3IiLCJoaWdobGlnaHRDb2xvciIsImJhY2tncm91bmRDb2xvciIsIm91dHB1dERpdiIsImlkIiwiZm9udFNpemUiLCJtYXJnaW5Ub3AiLCJoZWlnaHQiLCJjb2xvciIsImJvcmRlciIsInBhZGRpbmciLCJmb250RmFtaWx5IiwicmVzaXplIiwib3ZlcmZsb3dZIiwid2hpdGVTcGFjZSIsImxvZ29JbWFnZSIsImJvcmRlclJhZGl1cyIsImxvY2F0aW9uIiwiaHJlZiIsIl9hZGRUZXh0VG9EaXYiLCJ0ZXh0QXJlYSIsInF1ZXJ5U2VsZWN0b3IiLCJtYWtlTGFtYVJlcXVlc3QiLCJ0ZXh0Q29udGVudCIsIl94MyIsIl94NCIsIl94NSIsIl9tYWtlTGFtYVJlcXVlc3QiLCJjb21tYW5kIiwianNvbk9iamVjdCIsImZvcm1hdHRlZEJvZHkiLCJmb3JtYXR0ZWRKc29uIiwib3V0cHV0RmllbGQiLCJnb1dlYlJlcXVlc3RGdW5jIiwiSlNPTiIsInBhcnNlIiwiQm9keSIsInN0cmluZ2lmeSIsIl9jbGFzc0NhbGxDaGVjayIsIkNvbnN0cnVjdG9yIiwiX2RlZmluZVByb3BlcnRpZXMiLCJ0YXJnZXQiLCJwcm9wcyIsImRlc2NyaXB0b3IiLCJfdG9Qcm9wZXJ0eUtleSIsIl9jcmVhdGVDbGFzcyIsInByb3RvUHJvcHMiLCJzdGF0aWNQcm9wcyIsIl90b1ByaW1pdGl2ZSIsIlN0cmluZyIsImlucHV0IiwiaGludCIsInByaW0iLCJ0b1ByaW1pdGl2ZSIsInJlcyIsIk51bWJlciIsImVub3N5cyIsImNvZGUiLCJnbG9iYWxUaGlzIiwiZnMiLCJvdXRwdXRCdWYiLCJjb25zdGFudHMiLCJPX1dST05MWSIsIk9fUkRXUiIsIk9fQ1JFQVQiLCJPX1RSVU5DIiwiT19BUFBFTkQiLCJPX0VYQ0wiLCJ3cml0ZVN5bmMiLCJmZCIsImJ1ZiIsImRlY29kZXIiLCJkZWNvZGUiLCJubCIsImxhc3RJbmRleE9mIiwic3Vic3RyaW5nIiwid3JpdGUiLCJvZmZzZXQiLCJjYWxsYmFjayIsImNobW9kIiwicGF0aCIsIm1vZGUiLCJjaG93biIsInVpZCIsImdpZCIsImNsb3NlIiwiZmNobW9kIiwiZmNob3duIiwiZnN0YXQiLCJmc3luYyIsImZ0cnVuY2F0ZSIsImxjaG93biIsImxpbmsiLCJsc3RhdCIsIm1rZGlyIiwicGVybSIsIm9wZW4iLCJmbGFncyIsInJlYWQiLCJyZWFkZGlyIiwicmVhZGxpbmsiLCJyZW5hbWUiLCJmcm9tIiwidG8iLCJybWRpciIsInN0YXQiLCJzeW1saW5rIiwidHJ1bmNhdGUiLCJ1bmxpbmsiLCJ1dGltZXMiLCJhdGltZSIsIm10aW1lIiwicHJvY2VzcyIsImdldHVpZCIsImdldGdpZCIsImdldGV1aWQiLCJnZXRlZ2lkIiwiZ2V0Z3JvdXBzIiwicGlkIiwicHBpZCIsInVtYXNrIiwiY3dkIiwiY2hkaXIiLCJjcnlwdG8iLCJwZXJmb3JtYW5jZSIsIlRleHRFbmNvZGVyIiwiVGV4dERlY29kZXIiLCJlbmNvZGVyIiwiX2NsYXNzIiwiX3RoaXMiLCJhcmd2IiwiZW52IiwiZXhpdCIsIl9leGl0UHJvbWlzZSIsIl9yZXNvbHZlRXhpdFByb21pc2UiLCJfcGVuZGluZ0V2ZW50IiwiX3NjaGVkdWxlZFRpbWVvdXRzIiwiTWFwIiwiX25leHRDYWxsYmFja1RpbWVvdXRJRCIsInNldEludDY0IiwiYWRkciIsIm1lbSIsInNldFVpbnQzMiIsImZsb29yIiwic2V0SW50MzIiLCJnZXRJbnQ2NCIsImxvdyIsImdldFVpbnQzMiIsImhpZ2giLCJnZXRJbnQzMiIsImxvYWRWYWx1ZSIsImdldEZsb2F0NjQiLCJfdmFsdWVzIiwic3RvcmVWYWx1ZSIsIm5hbkhlYWQiLCJzZXRGbG9hdDY0IiwiX2lkcyIsImdldCIsIl9pZFBvb2wiLCJfZ29SZWZDb3VudHMiLCJzZXQiLCJ0eXBlRmxhZyIsImxvYWRTbGljZSIsImFycmF5IiwibGVuIiwiVWludDhBcnJheSIsIl9pbnN0IiwiZXhwb3J0cyIsImxvYWRTbGljZU9mVmFsdWVzIiwiQXJyYXkiLCJsb2FkU3RyaW5nIiwic2FkZHIiLCJEYXRhVmlldyIsInRpbWVPcmlnaW4iLCJEYXRlIiwibm93IiwiX2dvdGVzdCIsImFkZCIsImdvanMiLCJydW50aW1lV2FzbUV4aXQiLCJzcCIsImV4aXRlZCIsInJ1bnRpbWVXYXNtV3JpdGUiLCJydW50aW1lUmVzZXRNZW1vcnlEYXRhVmlldyIsInJ1bnRpbWVOYW5vdGltZTEiLCJydW50aW1lV2FsbHRpbWUiLCJtc2VjIiwiZ2V0VGltZSIsInJ1bnRpbWVTY2hlZHVsZVRpbWVvdXRFdmVudCIsInNldFRpbWVvdXQiLCJfcmVzdW1lIiwiaGFzIiwicnVudGltZUNsZWFyVGltZW91dEV2ZW50IiwiY2xlYXJUaW1lb3V0IiwicnVudGltZUdldFJhbmRvbURhdGEiLCJnZXRSYW5kb21WYWx1ZXMiLCJzeXNjYWxsSnNGaW5hbGl6ZVJlZiIsInN5c2NhbGxKc1N0cmluZ1ZhbCIsInN5c2NhbGxKc1ZhbHVlR2V0IiwiUmVmbGVjdCIsImdldHNwIiwic3lzY2FsbEpzVmFsdWVTZXQiLCJzeXNjYWxsSnNWYWx1ZURlbGV0ZSIsImRlbGV0ZVByb3BlcnR5Iiwic3lzY2FsbEpzVmFsdWVJbmRleCIsInN5c2NhbGxKc1ZhbHVlU2V0SW5kZXgiLCJzeXNjYWxsSnNWYWx1ZUNhbGwiLCJtIiwic2V0VWludDgiLCJzeXNjYWxsSnNWYWx1ZUludm9rZSIsInN5c2NhbGxKc1ZhbHVlTmV3IiwiY29uc3RydWN0Iiwic3lzY2FsbEpzVmFsdWVMZW5ndGgiLCJwYXJzZUludCIsInN5c2NhbGxKc1ZhbHVlUHJlcGFyZVN0cmluZyIsInN0ciIsImVuY29kZSIsInN5c2NhbGxKc1ZhbHVlTG9hZFN0cmluZyIsInN5c2NhbGxKc1ZhbHVlSW5zdGFuY2VPZiIsInN5c2NhbGxKc0NvcHlCeXRlc1RvR28iLCJkc3QiLCJVaW50OENsYW1wZWRBcnJheSIsInRvQ29weSIsInN1YmFycmF5Iiwic3lzY2FsbEpzQ29weUJ5dGVzVG9KUyIsImRlYnVnIiwiX3J1biIsIl90aGlzMiIsInN0clB0ciIsImFyZ2MiLCJhcmd2UHRycyIsIndhc21NaW5EYXRhQWRkciIsIkluc3RhbmNlIiwiTmFOIiwiZmlsbCIsIkluZmluaXR5IiwicHRyIiwiYnl0ZXMiLCJzb3J0IiwicmVzdW1lIiwiX21ha2VGdW5jV3JhcHBlciIsImV2ZW50Il0sIm1hcHBpbmdzIjoiO1FBQUE7UUFDQTs7UUFFQTtRQUNBOztRQUVBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBOztRQUVBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7OztRQUdBO1FBQ0E7O1FBRUE7UUFDQTs7UUFFQTtRQUNBO1FBQ0E7UUFDQTtRQUNBO1FBQ0E7UUFDQTtRQUNBLEtB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QUFBO2dCQUFBLE9BQUF0TCxRQUFBLENBQUE1QyxJQUFBO1lBQUE7VUFBQSxHQUFBd0MsT0FBQTtRQUFBLENBQ3ZCO1FBQUEsU0FBQTBCLElBQUE3QixFQUFBO1VBQUEsT0FBQTBSLElBQUEsQ0FBQTNTLEtBQUEsT0FBQUQsU0FBQTtRQUFBO1FBQUEsT0FBQStDLEdBQUE7TUFBQTtJQUFBO01BQUF0RCxHQUFBO01BQUFsRyxLQUFBLEVBRUQsU0FBQWlYLFFBQUEsRUFBVTtRQUNULElBQUksSUFBSSxDQUFDVCxNQUFNLEVBQUU7VUFDaEIsTUFBTSxJQUFJM1QsS0FBSyxDQUFDLCtCQUErQixDQUFDO1FBQ2pEO1FBQ0EsSUFBSSxDQUFDNFMsS0FBSyxDQUFDQyxPQUFPLENBQUN3RSxNQUFNLENBQUMsQ0FBQztRQUMzQixJQUFJLElBQUksQ0FBQzFELE1BQU0sRUFBRTtVQUNoQixJQUFJLENBQUMvQyxtQkFBbUIsQ0FBQyxDQUFDO1FBQzNCO01BQ0Q7SUFBQztNQUFBdk4sR0FBQTtNQUFBbEcsS0FBQSxFQUVELFNBQUFtYSxpQkFBaUI1TyxFQUFFLEVBQUU7UUFDcEIsSUFBTXpFLEVBQUUsR0FBRyxJQUFJO1FBQ2YsT0FBTyxZQUFZO1VBQ2xCLElBQU1zVCxLQUFLLEdBQUc7WUFBRTdPLEVBQUUsRUFBRUEsRUFBRTtZQUFFLFFBQU0sSUFBSTtZQUFFL0UsSUFBSSxFQUFFQztVQUFVLENBQUM7VUFDckRLLEVBQUUsQ0FBQzRNLGFBQWEsR0FBRzBHLEtBQUs7VUFDeEJ0VCxFQUFFLENBQUNtUSxPQUFPLENBQUMsQ0FBQztVQUNaLE9BQU9tRCxLQUFLLENBQUNyUixNQUFNO1FBQ3BCLENBQUM7TUFDRjtJQUFDO0lBQUEsT0FBQW9LLE1BQUE7RUFBQSxHQUNEO0FBQ0YsQ0FBQyxFQUFFLENBQUMsQyIsImZpbGUiOiJ3aWRnZXQuanMiLCJzb3VyY2VzQ29udGVudCI6WyIgXHQvLyBUaGUgbW9kdWxlIGNhY2hlXG4gXHR2YXIgaW5zdGFsbGVkTW9kdWxlcyA9IHt9O1xuXG4gXHQvLyBUaGUgcmVxdWlyZSBmdW5jdGlvblxuIFx0ZnVuY3Rpb24gX193ZWJwYWNrX3JlcXVpcmVfXyhtb2R1bGVJZCkge1xuXG4gXHRcdC8vIENoZWNrIGlmIG1vZHVsZSBpcyBpbiBjYWNoZVxuIFx0XHRpZihpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSkge1xuIFx0XHRcdHJldHVybiBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXS5leHBvcnRzO1xuIFx0XHR9XG4gXHRcdC8vIENyZWF0ZSBhIG5ldyBtb2R1bGUgKGFuZCBwdXQgaXQgaW50byB0aGUgY2FjaGUpXG4gXHRcdHZhciBtb2R1bGUgPSBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSA9IHtcbiBcdFx0XHRpOiBtb2R1bGVJZCxcbiBcdFx0XHRsOiBmYWxzZSxcbiBcdFx0XHRleHBvcnRzOiB7fVxuIFx0XHR9O1xuXG4gXHRcdC8vIEV4ZWN1dGUgdGhlIG1vZHVsZSBmdW5jdGlvblxuIFx0XHRtb2R1bGVzW21vZHVsZUlkXS5jYWxsKG1vZHVsZS5leHBvcnRzLCBtb2R1bGUsIG1vZHVsZS5leHBvcnRzLCBfX3dlYnBhY2tfcmVxdWlyZV9fKTtcblxuIFx0XHQvLyBGbGFnIHRoZSBtb2R1bGUgYXMgbG9hZGVkXG4gXHRcdG1vZHVsZS5sID0gdHJ1ZTtcblxuIFx0XHQvLyBSZXR1cm4gdGhlIGV4cG9ydHMgb2YgdGhlIG1vZHVsZVxuIFx0XHRyZXR1cm4gbW9kdWxlLmV4cG9ydHM7XG4gXHR9XG5cblxuIFx0Ly8gZXhwb3NlIHRoZSBtb2R1bGVzIG9iamVjdCAoX193ZWJwYWNrX21vZHVsZXNfXylcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubSA9IG1vZHVsZXM7XG5cbiBcdC8vIGV4cG9zZSB0aGUgbW9kdWxlIGNhY2hlXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLmMgPSBpbnN0YWxsZWRNb2R1bGVzO1xuXG4gXHQvLyBkZWZpbmUgZ2V0dGVyIGZ1bmN0aW9uIGZvciBoYXJtb255IGV4cG9ydHNcbiBcdF9fd2VicGFja19yZXF1aXJlX18uZCA9IGZ1bmN0aW9uKGV4cG9ydHMsIG5hbWUsIGdldHRlcikge1xuIFx0XHRpZighX193ZWJwYWNrX3JlcXVpcmVfXy5vKGV4cG9ydHMsIG5hbWUpKSB7XG4gXHRcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIG5hbWUsIHtcbiBcdFx0XHRcdGNvbmZpZ3VyYWJsZTogZmFsc2UsXG4gXHRcdFx0XHRlbnVtZXJhYmxlOiB0cnVlLFxuIFx0XHRcdFx0Z2V0OiBnZXR0ZXJcbiBcdFx0XHR9KTtcbiBcdFx0fVxuIFx0fTtcblxuIFx0Ly8gZ2V0RGVmYXVsdEV4cG9ydCBmdW5jdGlvbiBmb3IgY29tcGF0aWJpbGl0eSB3aXRoIG5vbi1oYXJtb255IG1vZHVsZXNcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubiA9IGZ1bmN0aW9uKG1vZHVsZSkge1xuIFx0XHR2YXIgZ2V0dGVyID0gbW9kdWxlICYmIG1vZHVsZS5fX2VzTW9kdWxlID9cbiBcdFx0XHRmdW5jdGlvbiBnZXREZWZhdWx0KCkgeyByZXR1cm4gbW9kdWxlWydkZWZhdWx0J107IH0gOlxuIFx0XHRcdGZ1bmN0aW9uIGdldE1vZHVsZUV4cG9ydHMoKSB7IHJldHVybiBtb2R1bGU7IH07XG4gXHRcdF9fd2VicGFja19yZXF1aXJlX18uZChnZXR0ZXIsICdhJywgZ2V0dGVyKTtcbiBcdFx0cmV0dXJuIGdldHRlcjtcbiBcdH07XG5cbiBcdC8vIE9iamVjdC5wcm90b3R5cGUuaGFzT3duUHJvcGVydHkuY2FsbFxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5vID0gZnVuY3Rpb24ob2JqZWN0LCBwcm9wZXJ0eSkgeyByZXR1cm4gT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eS5jYWxsKG9iamVjdCwgcHJvcGVydHkpOyB9O1xuXG4gXHQvLyBfX3dlYnBhY2tfcHVibGljX3BhdGhfX1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5wID0gXCJcIjtcblxuIFx0Ly8gTG9hZCBlbnRyeSBtb2R1bGUgYW5kIHJldHVybiBleHBvcnRzXG4gXHRyZXR1cm4gX193ZWJwYWNrX3JlcXVpcmVfXyhfX3dlYnBhY2tfcmVxdWlyZV9fLnMgPSAwKTtcblxuXG5cbi8vIFdFQlBBQ0sgRk9PVEVSIC8vXG4vLyB3ZWJwYWNrL2Jvb3RzdHJhcCAzM2M2ZjYxZGUxM2NiNzEwNzM3NCIsImltcG9ydCB7IHBpbmcgfSBmcm9tICcuL3NlcnZpY2VzJ1xuLy8gaW1wb3J0IHsgc2hvdyB9IGZyb20gJy4vdmlld3MvbWVzc2FnZSdcbmltcG9ydCB7IGFkZEJ1dHRvblRvVGV4dGFyZWEgfSBmcm9tICcuL3ZpZXdzL2xiJztcblxuY29uc3Qgc3VwcG9ydGVkQVBJID0gWydpbml0JywnbGFtYWJ1dHRvbiddOyAvLyBlbmxpc3QgYWxsIG1ldGhvZHMgc3VwcG9ydGVkIGJ5IEFQSSAoZS5nLiBgbXcoJ2V2ZW50JywgJ3VzZXItbG9naW4nKTtgKVxubGV0IGdvOyBcbi8qKlxuICAgIFRoZSBtYWluIGVudHJ5IG9mIHRoZSBhcHBsaWNhdGlvblxuICAgICovXG5mdW5jdGlvbiBhcHAod2luZG93KSB7XG4gICAgY29uc29sZS5sb2coJ0pTLVdpZGdldCBzdGFydGluZycpO1xuXG4gICAgLy8gc2V0IGRlZmF1bHQgY29uZmlndXJhdGlvbnNcbiAgICBsZXQgY29uZmlndXJhdGlvbnMgPSB7XG4gICAgICAgIHNvbWVEZWZhdWx0Q29uZmlndXJhdGlvbjogZmFsc2VcbiAgICB9O1xuXG4gICAgLy8gYWxsIG1ldGhvZHMgdGhhdCB3ZXJlIGNhbGxlZCB0aWxsIG5vdyBhbmQgc3RvcmVkIGluIHF1ZXVlXG4gICAgLy8gbmVlZHMgdG8gYmUgY2FsbGVkIG5vdyBcbiAgICBsZXQgZ2xvYmFsT2JqZWN0ID0gd2luZG93W3dpbmRvd1snSlMtV2lkZ2V0J11dO1xuICAgIGxldCBxdWV1ZSA9IGdsb2JhbE9iamVjdC5xO1xuICAgIGlmIChxdWV1ZSkge1xuICAgICAgICBmb3IgKHZhciBpID0gMDsgaSA8IHF1ZXVlLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBpZiAocXVldWVbaV1bMF0udG9Mb3dlckNhc2UoKSA9PSAnaW5pdCcpIHtcbiAgICAgICAgICAgICAgICBjb25maWd1cmF0aW9ucyA9IGV4dGVuZE9iamVjdChjb25maWd1cmF0aW9ucywgcXVldWVbaV1bMV0pO1xuICAgICAgICAgICAgICAgIGNvbnNvbGUubG9nKCdKUy1XaWRnZXQgc3RhcnRlZCcsIGNvbmZpZ3VyYXRpb25zKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGVsc2VcbiAgICAgICAgICAgICAgICBhcGlIYW5kbGVyKHF1ZXVlW2ldWzBdLCBxdWV1ZVtpXVsxXSk7XG4gICAgICAgIH1cbiAgICB9XG5cbiAgICAvLyBvdmVycmlkZSB0ZW1wb3JhcnkgKHVudGlsIHRoZSBhcHAgbG9hZGVkKSBoYW5kbGVyXG4gICAgLy8gZm9yIHdpZGdldCdzIEFQSSBjYWxsc1xuICAgIGdsb2JhbE9iamVjdCA9IGFwaUhhbmRsZXI7XG4gICAgZ2xvYmFsT2JqZWN0LmNvbmZpZ3VyYXRpb25zID0gY29uZmlndXJhdGlvbnM7XG59XG5cbi8qKlxuICAgIE1ldGhvZCB0aGF0IGhhbmRsZXMgYWxsIEFQSSBjYWxsc1xuICAgICovXG5hc3luYyBmdW5jdGlvbiBhcGlIYW5kbGVyKGFwaSwgcGFyYW1zKSB7XG4gICAgaWYgKCFhcGkpIHRocm93IEVycm9yKCdBUEkgbWV0aG9kIHJlcXVpcmVkJyk7XG4gICAgYXBpID0gYXBpLnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoc3VwcG9ydGVkQVBJLmluZGV4T2YoYXBpKSA9PT0gLTEpIHRocm93IEVycm9yKGBNZXRob2QgJHthcGl9IGlzIG5vdCBzdXBwb3J0ZWRgKTtcblxuICAgIGNvbnNvbGUubG9nKGBIYW5kbGluZyBBUEkgY2FsbCAke2FwaX1gLCBwYXJhbXMpO1xuXG4gICAgc3dpdGNoIChhcGkpIHtcbiAgICAgICAgLy8gVE9ETzogYWRkIEFQSSBpbXBsZW1lbnRhdGlvblxuICAgICAgICBjYXNlICdsYW1hYnV0dG9uJzpcbiAgICAgICAgICAgIC8vIGlmICghZ28pIHtcbiAgICAgICAgICAgICAgICBnbyA9IGF3YWl0IGxvYWRXYXNtKCk7IC8vIEVuc3VyZSBXYXNtIGlzIGxvYWRlZCBhbmQgR28gaW5zdGFuY2UgaXMgY3JlYXRlZFxuICAgICAgICAgICAgLy8gICAgIGlmICghZ28pIHJldHVybjsgLy8gRXhpdCBpZiB0aGVyZSB3YXMgYW4gZXJyb3IgbG9hZGluZyBXYXNtXG4gICAgICAgICAgICAvLyB9XG4gICAgICAgICAgICBhZGRCdXR0b25Ub1RleHRhcmVhKHBhcmFtcyxnbyk7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgZGVmYXVsdDpcbiAgICAgICAgICAgIGNvbnNvbGUud2FybihgTm8gaGFuZGxlciBkZWZpbmVkIGZvciAke2FwaX1gKTtcbiAgICB9XG59XG5cbmZ1bmN0aW9uIGV4dGVuZE9iamVjdChhLCBiKSB7XG4gICAgZm9yICh2YXIga2V5IGluIGIpXG4gICAgICAgIGlmIChiLmhhc093blByb3BlcnR5KGtleSkpXG4gICAgICAgICAgICBhW2tleV0gPSBiW2tleV07XG4gICAgcmV0dXJuIGE7XG59XG5cbmFwcCh3aW5kb3cpO1xuXG5hc3luYyBmdW5jdGlvbiBsb2FkV2FzbSgpIHtcbiAgICB0cnkge1xuICAgICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKFwibWFpbi53YXNtXCIpO1xuICAgICAgICBjb25zdCBidWZmZXIgPSBhd2FpdCByZXNwb25zZS5hcnJheUJ1ZmZlcigpO1xuICAgICAgICBjb25zdCBnbyA9IG5ldyBHbygpO1xuICAgICAgICBjb25zdCByZXN1bHQgPSBhd2FpdCBXZWJBc3NlbWJseS5pbnN0YW50aWF0ZShidWZmZXIsIGdvLmltcG9ydE9iamVjdCk7XG4gICAgICAgIGdvLnJ1bihyZXN1bHQuaW5zdGFuY2UpO1xuICAgICAgICBjb25zb2xlLmxvZyhcIldhc20gbG9hZGVkIGFuZCBHbyBpbnN0YW5jZSBpbml0aWFsaXplZFwiKTtcbiAgICAgICAgcmV0dXJuIGdvO1xuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoXCJFcnJvciBsb2FkaW5nIFdhc206XCIsIGVycm9yKTtcbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxufVxuXG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL21haW4uanMiLCJcbmV4cG9ydCBmdW5jdGlvbiBwaW5nKCkge1xuICAgIHJldHVybiAncG9uZyc7XG59XG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL3NlcnZpY2VzLmpzIiwiaW1wb3J0ICcuL2xiLmNzcyc7XG5pbXBvcnQgJy4vd2FzbV9leGVjJ1xuXG5leHBvcnQgZnVuY3Rpb24gYWRkQnV0dG9uVG9UZXh0YXJlYShjbGFzc05hbWUsZ28pIHtcbiAgICBpZiAoIWdvKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKFwid2FzbSBmaWxlIG5vdCBpbml0aWFsaXplZFwiKVxuICAgICAgICAgICAgcmV0dXJuOyBcbiAgICAgICAgfVxuICAgIC8vIGNvbnN0IGRpdkVsZW1lbnRzID0gZG9jdW1lbnQucXVlcnlTZWxlY3RvckFsbChgZGl2LiR7Y2xhc3NOYW1lfWApO1xuICAgIC8vIFNlbGVjdCBhbGwgZWxlbWVudHMgd2l0aCB0aGUgY2xhc3MgbmFtZSAnbXktYnV0dG9uLWNvbnRhaW5lcidcbiAgICBjb25zdCBkaXZFbGVtZW50cyA9IGRvY3VtZW50LnF1ZXJ5U2VsZWN0b3JBbGwoYC4ke2NsYXNzTmFtZX1gKTtcblxuICAgIGRpdkVsZW1lbnRzLmZvckVhY2goKGRpdkVsZW1lbnQpID0+IHtcblxuICAgICAgICBjb25zdCBidXR0b25JbWFnZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2ltZycpO1xuICAgICAgICBidXR0b25JbWFnZS5zcmMgPSAnYXJyb3dsMi5wbmcnO1xuICAgICAgICBidXR0b25JbWFnZS5zdHlsZS53aWR0aCA9ICczMHB4JzsgXG4gICAgICAgIGNvbnN0IGluZGV4ID0gTWF0aC5yYW5kb20oKS50b1N0cmluZygzNikuc3Vic3RyKDIsIDEwKTtcbiAgICAgICAgY29uc3QgYnV0dG9uRWxlbWVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2J1dHRvbicpO1xuICAgICAgICBidXR0b25FbGVtZW50LmFwcGVuZENoaWxkKGJ1dHRvbkltYWdlKTtcbiAgICAgICAgYnV0dG9uRWxlbWVudC5hZGRFdmVudExpc3RlbmVyKCdjbGljaycsICgpID0+IGFkZFRleHRUb0RpdihkaXZFbGVtZW50LGluZGV4KSk7XG4gICAgICAgIGJ1dHRvbkVsZW1lbnQuc3R5bGUuYWxsID0gJ3Vuc2V0JztcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5wb3NpdGlvbiA9ICdyZWxhdGl2ZSc7XG4gICAgICAgIGJ1dHRvbkVsZW1lbnQuc3R5bGUuYm90dG9tID0gJzE1MHB4JztcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5yaWdodCA9ICc1MHB4JztcblxuXG4gICAgICAgIC8vIEFkZCBob3ZlciBlZmZlY3RcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS50cmFuc2l0aW9uID0gJ2JhY2tncm91bmQtY29sb3IgMC4zcyc7XG4gICAgICAgIGJ1dHRvbkVsZW1lbnQuc3R5bGUuY3Vyc29yID0gJ3BvaW50ZXInO1xuICAgICAgICBjb25zdCBoaWdobGlnaHRDb2xvciA9ICcjMDBBMUUxJztcbiAgICAgICAgYnV0dG9uRWxlbWVudC5hZGRFdmVudExpc3RlbmVyKCdtb3VzZWVudGVyJywgKCkgPT4ge1xuICAgICAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSBoaWdobGlnaHRDb2xvcjsgLy8gQ2hhbmdlIHRoZSBiYWNrZ3JvdW5kIGNvbG9yIG9uIGhvdmVyXG4gICAgICAgICAgfSk7XG4gICAgICAgICAgXG4gICAgICAgICAgYnV0dG9uRWxlbWVudC5hZGRFdmVudExpc3RlbmVyKCdtb3VzZWxlYXZlJywgKCkgPT4ge1xuICAgICAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSAnJzsgLy8gUmVzZXQgdGhlIGJhY2tncm91bmQgY29sb3Igb24gbW91c2UgbGVhdmVcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBcbiAgICAgICAgICBidXR0b25FbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNlZG93bicsICgpID0+IHtcbiAgICAgICAgICAgIGJ1dHRvbkVsZW1lbnQuc3R5bGUuYmFja2dyb3VuZENvbG9yID0gaGlnaGxpZ2h0Q29sb3I7IC8vIENoYW5nZSB0aGUgYmFja2dyb3VuZCBjb2xvciB3aGVuIGNsaWNrZWRcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBcbiAgICAgICAgICBidXR0b25FbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNldXAnLCAoKSA9PiB7XG4gICAgICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLmJhY2tncm91bmRDb2xvciA9IGhpZ2hsaWdodENvbG9yOyAvLyBSZXN0b3JlIHRoZSBob3ZlciBjb2xvciBvbiByZWxlYXNlXG4gICAgICAgICAgfSk7XG4gICAgICAgICAgXG5cbiAgICAgICAgZGl2RWxlbWVudC5hcHBlbmRDaGlsZChidXR0b25FbGVtZW50KTtcblxuICAgICAgICBjb25zdCBvdXRwdXREaXYgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KCdkaXYnKTtcbiAgICAgICAgb3V0cHV0RGl2LmNsYXNzID0gJ291dHB1dEZpZWxkJ1xuICAgICAgICBvdXRwdXREaXYuaWQgPSBgb3V0cHV0RmllbGRfJHtpbmRleH1gO1xuICAgICAgICBvdXRwdXREaXYuc3R5bGUuZm9udFNpemUgPSAnMTZweCc7XG4gICAgICAgIG91dHB1dERpdi5zdHlsZS5tYXJnaW5Ub3AgPSAnMTBweCc7XG4gICAgICAgIG91dHB1dERpdi5zdHlsZS53aWR0aCA9ICc1MDBweCc7XG4gICAgICAgIG91dHB1dERpdi5zdHlsZS5oZWlnaHQgPSAnMzAwcHgnO1xuICAgICAgICBvdXRwdXREaXYuc3R5bGUuY29sb3IgPSAnd2hpdGUnO1xuICAgICAgICBvdXRwdXREaXYuc3R5bGUuYmFja2dyb3VuZENvbG9yID0gJ2JsYWNrJztcbiAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmJvcmRlciA9ICcxcHggc29saWQgIzU1NSc7XG4gICAgICAgIG91dHB1dERpdi5zdHlsZS5wYWRkaW5nID0gJzEwcHgnO1xuICAgICAgICBvdXRwdXREaXYuc3R5bGUuZm9udEZhbWlseSA9ICdtb25vc3BhY2UnO1xuICAgICAgICBvdXRwdXREaXYuc3R5bGUucmVzaXplID0gJ3ZlcnRpY2FsJztcbiAgICAgICAgb3V0cHV0RGl2LnN0eWxlLm92ZXJmbG93WSA9ICdhdXRvJztcbiAgICAgICAgb3V0cHV0RGl2LnN0eWxlLndoaXRlU3BhY2UgPSAncHJlLXdyYXAnO1xuXG4gICAgICAgIGNvbnN0IGxvZ29JbWFnZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2ltZycpO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUucG9zaXRpb24gPSAncmVsYXRpdmUnO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUuYm90dG9tID0gJzEwcHgnO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUucmlnaHQgPSAnODBweCc7XG4gICAgICAgIGxvZ29JbWFnZS5zcmMgPSAnaGV4bW9zX2xvZ28uanBlZyc7XG4gICAgICAgIGxvZ29JbWFnZS5zdHlsZS53aWR0aCA9ICczMHB4JztcbiAgICAgICAgbG9nb0ltYWdlLnN0eWxlLmJvcmRlclJhZGl1cyA9ICc1MCUnOyBcblxuICAgICAgICBsb2dvSW1hZ2UuYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCBmdW5jdGlvbigpIHtcbiAgICAgICAgICAgIHdpbmRvdy5sb2NhdGlvbi5ocmVmID0gJ2h0dHBzOi8vaGV4bW9zLmNvbSc7XG4gICAgICAgICAgfSk7XG4gICAgICAgICAgXG4gICAgICAgIGRpdkVsZW1lbnQuYXBwZW5kQ2hpbGQobG9nb0ltYWdlKVxuXG5cbiAgICAgICAgZGl2RWxlbWVudC5hcHBlbmRDaGlsZChvdXRwdXREaXYpO1xuICAgIH0pO1xufVxuXG5hc3luYyBmdW5jdGlvbiBhZGRUZXh0VG9EaXYoZGl2RWxlbWVudCxpbmRleCkge1xuICAgIGNvbnN0IHRleHRBcmVhID0gZGl2RWxlbWVudC5xdWVyeVNlbGVjdG9yKCd0ZXh0YXJlYScpO1xuICAgIGlmICh0ZXh0QXJlYSkge1xuICAgICAgICBjb25zb2xlLmxvZyhcIlRleHQgaW5zaWRlIHRoZSB0ZXh0IGFyZWE6XCIsIHRleHRBcmVhLnZhbHVlKTtcbiAgICAgICAgbWFrZUxhbWFSZXF1ZXN0KHRleHRBcmVhLnZhbHVlLCBkaXZFbGVtZW50LCBpbmRleClcbiAgICB9XG4gICAgZWxzZXtcbiAgICAgICAgY29uc29sZS5sb2coXCJObyB0ZXh0IGFyZWEgZm91bmQuIExvZ2dpbmcgZGF0YSBmcm9tIGRpdkVsZW1lbnQ6XCIsIGRpdkVsZW1lbnQudGV4dENvbnRlbnQpO1xuICAgICAgICBtYWtlTGFtYVJlcXVlc3QoZGl2RWxlbWVudC50ZXh0Q29udGVudCxkaXZFbGVtZW50LCBpbmRleClcbiAgICB9XG59XG5cbmFzeW5jIGZ1bmN0aW9uIG1ha2VMYW1hUmVxdWVzdChjb21tYW5kLCBkaXZFbGVtZW50LCBpbmRleCkge1xuICAgIHRyeSB7XG4gICAgICAgIGlmICghZ28pIHtcbiAgICAgICAgICAgIGNvbnNvbGUuZXJyb3IoXCJXYXNtIGlzIG5vdCBsb2FkZWQuIENhbm5vdCBtYWtlIGEgcmVxdWVzdC5cIik7XG4gICAgICAgICAgICByZXR1cm47XG4gICAgICAgIH1cbiAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGdvV2ViUmVxdWVzdEZ1bmMoY29tbWFuZCk7XG4gICAgY29uc29sZS5sb2coXCJyZXNwb25zZSBEYXRhOlwiLHJlc3BvbnNlKVxuICAgIGNvbnN0IGpzb25PYmplY3QgPSBKU09OLnBhcnNlKHJlc3BvbnNlKTtcbiAgICBjb25zdCBmb3JtYXR0ZWRCb2R5ID0gSlNPTi5wYXJzZShqc29uT2JqZWN0LkJvZHkpOyBcbiAgICB2YXIgZm9ybWF0dGVkSnNvbiA9IEpTT04uc3RyaW5naWZ5KGZvcm1hdHRlZEJvZHksIG51bGwsIDIpO1xuICAgIGNvbnNvbGUubG9nKFwiZm9ybWF0dGVkSnNvbjpcIixmb3JtYXR0ZWRKc29uKSAgICBcbiAgICBjb25zdCBvdXRwdXRGaWVsZCA9IGRpdkVsZW1lbnQucXVlcnlTZWxlY3RvcihgI291dHB1dEZpZWxkXyR7aW5kZXh9YCk7XG4gICAgY29uc29sZS5sb2cob3V0cHV0RmllbGQpXG4gICAgb3V0cHV0RmllbGQudGV4dENvbnRlbnQgPSBmb3JtYXR0ZWRKc29uOyBcbiAgICBjb25zb2xlLmxvZyhcIm91dHB1dEZpZWxkLnRleHRDb250ZW50OlwiLG91dHB1dEZpZWxkLnRleHRDb250ZW50KVxufSBjYXRjaCAoZXJyb3IpIHtcbiAgICBjb25zb2xlLmVycm9yKFwiRXJyb3IgbWFraW5nIExhbWEgcmVxdWVzdDpcIiwgZXJyb3IpO1xufVxufVxuXG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL3ZpZXdzL2xiLmpzIiwiLy8gc3R5bGUtbG9hZGVyOiBBZGRzIHNvbWUgY3NzIHRvIHRoZSBET00gYnkgYWRkaW5nIGEgPHN0eWxlPiB0YWdcblxuLy8gbG9hZCB0aGUgc3R5bGVzXG52YXIgY29udGVudCA9IHJlcXVpcmUoXCIhIS4uLy4uL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvY2pzLmpzIS4vbGIuY3NzXCIpO1xuaWYodHlwZW9mIGNvbnRlbnQgPT09ICdzdHJpbmcnKSBjb250ZW50ID0gW1ttb2R1bGUuaWQsIGNvbnRlbnQsICcnXV07XG4vLyBQcmVwYXJlIGNzc1RyYW5zZm9ybWF0aW9uXG52YXIgdHJhbnNmb3JtO1xuXG52YXIgb3B0aW9ucyA9IHtcImhtclwiOnRydWV9XG5vcHRpb25zLnRyYW5zZm9ybSA9IHRyYW5zZm9ybVxuLy8gYWRkIHRoZSBzdHlsZXMgdG8gdGhlIERPTVxudmFyIHVwZGF0ZSA9IHJlcXVpcmUoXCIhLi4vLi4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9saWIvYWRkU3R5bGVzLmpzXCIpKGNvbnRlbnQsIG9wdGlvbnMpO1xuaWYoY29udGVudC5sb2NhbHMpIG1vZHVsZS5leHBvcnRzID0gY29udGVudC5sb2NhbHM7XG4vLyBIb3QgTW9kdWxlIFJlcGxhY2VtZW50XG5pZihtb2R1bGUuaG90KSB7XG5cdC8vIFdoZW4gdGhlIHN0eWxlcyBjaGFuZ2UsIHVwZGF0ZSB0aGUgPHN0eWxlPiB0YWdzXG5cdGlmKCFjb250ZW50LmxvY2Fscykge1xuXHRcdG1vZHVsZS5ob3QuYWNjZXB0KFwiISEuLi8uLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L2Nqcy5qcyEuL2xiLmNzc1wiLCBmdW5jdGlvbigpIHtcblx0XHRcdHZhciBuZXdDb250ZW50ID0gcmVxdWlyZShcIiEhLi4vLi4vbm9kZV9tb2R1bGVzL2Nzcy1sb2FkZXIvZGlzdC9janMuanMhLi9sYi5jc3NcIik7XG5cdFx0XHRpZih0eXBlb2YgbmV3Q29udGVudCA9PT0gJ3N0cmluZycpIG5ld0NvbnRlbnQgPSBbW21vZHVsZS5pZCwgbmV3Q29udGVudCwgJyddXTtcblx0XHRcdHVwZGF0ZShuZXdDb250ZW50KTtcblx0XHR9KTtcblx0fVxuXHQvLyBXaGVuIHRoZSBtb2R1bGUgaXMgZGlzcG9zZWQsIHJlbW92ZSB0aGUgPHN0eWxlPiB0YWdzXG5cdG1vZHVsZS5ob3QuZGlzcG9zZShmdW5jdGlvbigpIHsgdXBkYXRlKCk7IH0pO1xufVxuXG5cbi8vLy8vLy8vLy8vLy8vLy8vL1xuLy8gV0VCUEFDSyBGT09URVJcbi8vIC4vc3JjL3ZpZXdzL2xiLmNzc1xuLy8gbW9kdWxlIGlkID0gM1xuLy8gbW9kdWxlIGNodW5rcyA9IDAiLCIvLyBJbXBvcnRzXG52YXIgX19fQ1NTX0xPQURFUl9BUElfSU1QT1JUX19fID0gcmVxdWlyZShcIi4uLy4uL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvcnVudGltZS9hcGkuanNcIik7XG5leHBvcnRzID0gX19fQ1NTX0xPQURFUl9BUElfSU1QT1JUX19fKGZhbHNlKTtcbi8vIE1vZHVsZVxuZXhwb3J0cy5wdXNoKFttb2R1bGUuaWQsIFwiLyogbGFtYWJ1dHRvbi5jc3MgKi9cXG5cXG4vKiBTdHlsZSBmb3IgdGhlIGJ1dHRvbiBlbGVtZW50ICovXFxuLmpzLXdpZGdldC1sYW1hYnV0dG9uIHtcXG4gICAgYmFja2dyb3VuZC1jb2xvcjogZ3JlZW47IC8qIFNldCB0aGUgYmFja2dyb3VuZCBjb2xvciB0byBncmVlbiAqL1xcbiAgICBjb2xvcjogd2hpdGU7IC8qIFNldCB0aGUgdGV4dCBjb2xvciB0byB3aGl0ZSAqL1xcbiAgICBib3JkZXI6IG5vbmU7IC8qIFJlbW92ZSBib3JkZXJzICovXFxuICAgIHBhZGRpbmc6IDVweCAxMHB4OyAvKiBBZGp1c3QgcGFkZGluZyBmb3IgYSBzbWFsbGVyIGJ1dHRvbiAqL1xcbiAgICBjdXJzb3I6IHBvaW50ZXI7IC8qIEFkZCBhIHBvaW50ZXIgY3Vyc29yIG9uIGhvdmVyICovXFxuICAgIHBvc2l0aW9uOiBhYnNvbHV0ZTtcXG59XFxuXFxuLyogSG92ZXIgZWZmZWN0IHdoZW4gdGhlIGJ1dHRvbiBpcyBob3ZlcmVkIG92ZXIgKi9cXG4uanMtd2lkZ2V0LWxhbWFidXR0b246aG92ZXIge1xcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiBkYXJrZ3JlZW47IC8qIERhcmtlbiB0aGUgYmFja2dyb3VuZCBjb2xvciBvbiBob3ZlciAqL1xcbn1cXG5cIiwgXCJcIl0pO1xuLy8gRXhwb3J0c1xubW9kdWxlLmV4cG9ydHMgPSBleHBvcnRzO1xuXG5cblxuLy8vLy8vLy8vLy8vLy8vLy8vXG4vLyBXRUJQQUNLIEZPT1RFUlxuLy8gLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L2Nqcy5qcyEuL3NyYy92aWV3cy9sYi5jc3Ncbi8vIG1vZHVsZSBpZCA9IDRcbi8vIG1vZHVsZSBjaHVua3MgPSAwIiwiXCJ1c2Ugc3RyaWN0XCI7XG5cbi8qXG4gIE1JVCBMaWNlbnNlIGh0dHA6Ly93d3cub3BlbnNvdXJjZS5vcmcvbGljZW5zZXMvbWl0LWxpY2Vuc2UucGhwXG4gIEF1dGhvciBUb2JpYXMgS29wcGVycyBAc29rcmFcbiovXG4vLyBjc3MgYmFzZSBjb2RlLCBpbmplY3RlZCBieSB0aGUgY3NzLWxvYWRlclxuLy8gZXNsaW50LWRpc2FibGUtbmV4dC1saW5lIGZ1bmMtbmFtZXNcbm1vZHVsZS5leHBvcnRzID0gZnVuY3Rpb24gKHVzZVNvdXJjZU1hcCkge1xuICB2YXIgbGlzdCA9IFtdOyAvLyByZXR1cm4gdGhlIGxpc3Qgb2YgbW9kdWxlcyBhcyBjc3Mgc3RyaW5nXG5cbiAgbGlzdC50b1N0cmluZyA9IGZ1bmN0aW9uIHRvU3RyaW5nKCkge1xuICAgIHJldHVybiB0aGlzLm1hcChmdW5jdGlvbiAoaXRlbSkge1xuICAgICAgdmFyIGNvbnRlbnQgPSBjc3NXaXRoTWFwcGluZ1RvU3RyaW5nKGl0ZW0sIHVzZVNvdXJjZU1hcCk7XG5cbiAgICAgIGlmIChpdGVtWzJdKSB7XG4gICAgICAgIHJldHVybiBcIkBtZWRpYSBcIi5jb25jYXQoaXRlbVsyXSwgXCIge1wiKS5jb25jYXQoY29udGVudCwgXCJ9XCIpO1xuICAgICAgfVxuXG4gICAgICByZXR1cm4gY29udGVudDtcbiAgICB9KS5qb2luKCcnKTtcbiAgfTsgLy8gaW1wb3J0IGEgbGlzdCBvZiBtb2R1bGVzIGludG8gdGhlIGxpc3RcbiAgLy8gZXNsaW50LWRpc2FibGUtbmV4dC1saW5lIGZ1bmMtbmFtZXNcblxuXG4gIGxpc3QuaSA9IGZ1bmN0aW9uIChtb2R1bGVzLCBtZWRpYVF1ZXJ5LCBkZWR1cGUpIHtcbiAgICBpZiAodHlwZW9mIG1vZHVsZXMgPT09ICdzdHJpbmcnKSB7XG4gICAgICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgbm8tcGFyYW0tcmVhc3NpZ25cbiAgICAgIG1vZHVsZXMgPSBbW251bGwsIG1vZHVsZXMsICcnXV07XG4gICAgfVxuXG4gICAgdmFyIGFscmVhZHlJbXBvcnRlZE1vZHVsZXMgPSB7fTtcblxuICAgIGlmIChkZWR1cGUpIHtcbiAgICAgIGZvciAodmFyIGkgPSAwOyBpIDwgdGhpcy5sZW5ndGg7IGkrKykge1xuICAgICAgICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgcHJlZmVyLWRlc3RydWN0dXJpbmdcbiAgICAgICAgdmFyIGlkID0gdGhpc1tpXVswXTtcblxuICAgICAgICBpZiAoaWQgIT0gbnVsbCkge1xuICAgICAgICAgIGFscmVhZHlJbXBvcnRlZE1vZHVsZXNbaWRdID0gdHJ1ZTtcbiAgICAgICAgfVxuICAgICAgfVxuICAgIH1cblxuICAgIGZvciAodmFyIF9pID0gMDsgX2kgPCBtb2R1bGVzLmxlbmd0aDsgX2krKykge1xuICAgICAgdmFyIGl0ZW0gPSBbXS5jb25jYXQobW9kdWxlc1tfaV0pO1xuXG4gICAgICBpZiAoZGVkdXBlICYmIGFscmVhZHlJbXBvcnRlZE1vZHVsZXNbaXRlbVswXV0pIHtcbiAgICAgICAgLy8gZXNsaW50LWRpc2FibGUtbmV4dC1saW5lIG5vLWNvbnRpbnVlXG4gICAgICAgIGNvbnRpbnVlO1xuICAgICAgfVxuXG4gICAgICBpZiAobWVkaWFRdWVyeSkge1xuICAgICAgICBpZiAoIWl0ZW1bMl0pIHtcbiAgICAgICAgICBpdGVtWzJdID0gbWVkaWFRdWVyeTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICBpdGVtWzJdID0gXCJcIi5jb25jYXQobWVkaWFRdWVyeSwgXCIgYW5kIFwiKS5jb25jYXQoaXRlbVsyXSk7XG4gICAgICAgIH1cbiAgICAgIH1cblxuICAgICAgbGlzdC5wdXNoKGl0ZW0pO1xuICAgIH1cbiAgfTtcblxuICByZXR1cm4gbGlzdDtcbn07XG5cbmZ1bmN0aW9uIGNzc1dpdGhNYXBwaW5nVG9TdHJpbmcoaXRlbSwgdXNlU291cmNlTWFwKSB7XG4gIHZhciBjb250ZW50ID0gaXRlbVsxXSB8fCAnJzsgLy8gZXNsaW50LWRpc2FibGUtbmV4dC1saW5lIHByZWZlci1kZXN0cnVjdHVyaW5nXG5cbiAgdmFyIGNzc01hcHBpbmcgPSBpdGVtWzNdO1xuXG4gIGlmICghY3NzTWFwcGluZykge1xuICAgIHJldHVybiBjb250ZW50O1xuICB9XG5cbiAgaWYgKHVzZVNvdXJjZU1hcCAmJiB0eXBlb2YgYnRvYSA9PT0gJ2Z1bmN0aW9uJykge1xuICAgIHZhciBzb3VyY2VNYXBwaW5nID0gdG9Db21tZW50KGNzc01hcHBpbmcpO1xuICAgIHZhciBzb3VyY2VVUkxzID0gY3NzTWFwcGluZy5zb3VyY2VzLm1hcChmdW5jdGlvbiAoc291cmNlKSB7XG4gICAgICByZXR1cm4gXCIvKiMgc291cmNlVVJMPVwiLmNvbmNhdChjc3NNYXBwaW5nLnNvdXJjZVJvb3QgfHwgJycpLmNvbmNhdChzb3VyY2UsIFwiICovXCIpO1xuICAgIH0pO1xuICAgIHJldHVybiBbY29udGVudF0uY29uY2F0KHNvdXJjZVVSTHMpLmNvbmNhdChbc291cmNlTWFwcGluZ10pLmpvaW4oJ1xcbicpO1xuICB9XG5cbiAgcmV0dXJuIFtjb250ZW50XS5qb2luKCdcXG4nKTtcbn0gLy8gQWRhcHRlZCBmcm9tIGNvbnZlcnQtc291cmNlLW1hcCAoTUlUKVxuXG5cbmZ1bmN0aW9uIHRvQ29tbWVudChzb3VyY2VNYXApIHtcbiAgLy8gZXNsaW50LWRpc2FibGUtbmV4dC1saW5lIG5vLXVuZGVmXG4gIHZhciBiYXNlNjQgPSBidG9hKHVuZXNjYXBlKGVuY29kZVVSSUNvbXBvbmVudChKU09OLnN0cmluZ2lmeShzb3VyY2VNYXApKSkpO1xuICB2YXIgZGF0YSA9IFwic291cmNlTWFwcGluZ1VSTD1kYXRhOmFwcGxpY2F0aW9uL2pzb247Y2hhcnNldD11dGYtODtiYXNlNjQsXCIuY29uY2F0KGJhc2U2NCk7XG4gIHJldHVybiBcIi8qIyBcIi5jb25jYXQoZGF0YSwgXCIgKi9cIik7XG59XG5cblxuLy8vLy8vLy8vLy8vLy8vLy8vXG4vLyBXRUJQQUNLIEZPT1RFUlxuLy8gLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L3J1bnRpbWUvYXBpLmpzXG4vLyBtb2R1bGUgaWQgPSA1XG4vLyBtb2R1bGUgY2h1bmtzID0gMCIsIi8qXG5cdE1JVCBMaWNlbnNlIGh0dHA6Ly93d3cub3BlbnNvdXJjZS5vcmcvbGljZW5zZXMvbWl0LWxpY2Vuc2UucGhwXG5cdEF1dGhvciBUb2JpYXMgS29wcGVycyBAc29rcmFcbiovXG5cbnZhciBzdHlsZXNJbkRvbSA9IHt9O1xuXG52YXJcdG1lbW9pemUgPSBmdW5jdGlvbiAoZm4pIHtcblx0dmFyIG1lbW87XG5cblx0cmV0dXJuIGZ1bmN0aW9uICgpIHtcblx0XHRpZiAodHlwZW9mIG1lbW8gPT09IFwidW5kZWZpbmVkXCIpIG1lbW8gPSBmbi5hcHBseSh0aGlzLCBhcmd1bWVudHMpO1xuXHRcdHJldHVybiBtZW1vO1xuXHR9O1xufTtcblxudmFyIGlzT2xkSUUgPSBtZW1vaXplKGZ1bmN0aW9uICgpIHtcblx0Ly8gVGVzdCBmb3IgSUUgPD0gOSBhcyBwcm9wb3NlZCBieSBCcm93c2VyaGFja3Ncblx0Ly8gQHNlZSBodHRwOi8vYnJvd3NlcmhhY2tzLmNvbS8jaGFjay1lNzFkODY5MmY2NTMzNDE3M2ZlZTcxNWMyMjJjYjgwNVxuXHQvLyBUZXN0cyBmb3IgZXhpc3RlbmNlIG9mIHN0YW5kYXJkIGdsb2JhbHMgaXMgdG8gYWxsb3cgc3R5bGUtbG9hZGVyXG5cdC8vIHRvIG9wZXJhdGUgY29ycmVjdGx5IGludG8gbm9uLXN0YW5kYXJkIGVudmlyb25tZW50c1xuXHQvLyBAc2VlIGh0dHBzOi8vZ2l0aHViLmNvbS93ZWJwYWNrLWNvbnRyaWIvc3R5bGUtbG9hZGVyL2lzc3Vlcy8xNzdcblx0cmV0dXJuIHdpbmRvdyAmJiBkb2N1bWVudCAmJiBkb2N1bWVudC5hbGwgJiYgIXdpbmRvdy5hdG9iO1xufSk7XG5cbnZhciBnZXRFbGVtZW50ID0gKGZ1bmN0aW9uIChmbikge1xuXHR2YXIgbWVtbyA9IHt9O1xuXG5cdHJldHVybiBmdW5jdGlvbihzZWxlY3Rvcikge1xuXHRcdGlmICh0eXBlb2YgbWVtb1tzZWxlY3Rvcl0gPT09IFwidW5kZWZpbmVkXCIpIHtcblx0XHRcdHZhciBzdHlsZVRhcmdldCA9IGZuLmNhbGwodGhpcywgc2VsZWN0b3IpO1xuXHRcdFx0Ly8gU3BlY2lhbCBjYXNlIHRvIHJldHVybiBoZWFkIG9mIGlmcmFtZSBpbnN0ZWFkIG9mIGlmcmFtZSBpdHNlbGZcblx0XHRcdGlmIChzdHlsZVRhcmdldCBpbnN0YW5jZW9mIHdpbmRvdy5IVE1MSUZyYW1lRWxlbWVudCkge1xuXHRcdFx0XHR0cnkge1xuXHRcdFx0XHRcdC8vIFRoaXMgd2lsbCB0aHJvdyBhbiBleGNlcHRpb24gaWYgYWNjZXNzIHRvIGlmcmFtZSBpcyBibG9ja2VkXG5cdFx0XHRcdFx0Ly8gZHVlIHRvIGNyb3NzLW9yaWdpbiByZXN0cmljdGlvbnNcblx0XHRcdFx0XHRzdHlsZVRhcmdldCA9IHN0eWxlVGFyZ2V0LmNvbnRlbnREb2N1bWVudC5oZWFkO1xuXHRcdFx0XHR9IGNhdGNoKGUpIHtcblx0XHRcdFx0XHRzdHlsZVRhcmdldCA9IG51bGw7XG5cdFx0XHRcdH1cblx0XHRcdH1cblx0XHRcdG1lbW9bc2VsZWN0b3JdID0gc3R5bGVUYXJnZXQ7XG5cdFx0fVxuXHRcdHJldHVybiBtZW1vW3NlbGVjdG9yXVxuXHR9O1xufSkoZnVuY3Rpb24gKHRhcmdldCkge1xuXHRyZXR1cm4gZG9jdW1lbnQucXVlcnlTZWxlY3Rvcih0YXJnZXQpXG59KTtcblxudmFyIHNpbmdsZXRvbiA9IG51bGw7XG52YXJcdHNpbmdsZXRvbkNvdW50ZXIgPSAwO1xudmFyXHRzdHlsZXNJbnNlcnRlZEF0VG9wID0gW107XG5cbnZhclx0Zml4VXJscyA9IHJlcXVpcmUoXCIuL3VybHNcIik7XG5cbm1vZHVsZS5leHBvcnRzID0gZnVuY3Rpb24obGlzdCwgb3B0aW9ucykge1xuXHRpZiAodHlwZW9mIERFQlVHICE9PSBcInVuZGVmaW5lZFwiICYmIERFQlVHKSB7XG5cdFx0aWYgKHR5cGVvZiBkb2N1bWVudCAhPT0gXCJvYmplY3RcIikgdGhyb3cgbmV3IEVycm9yKFwiVGhlIHN0eWxlLWxvYWRlciBjYW5ub3QgYmUgdXNlZCBpbiBhIG5vbi1icm93c2VyIGVudmlyb25tZW50XCIpO1xuXHR9XG5cblx0b3B0aW9ucyA9IG9wdGlvbnMgfHwge307XG5cblx0b3B0aW9ucy5hdHRycyA9IHR5cGVvZiBvcHRpb25zLmF0dHJzID09PSBcIm9iamVjdFwiID8gb3B0aW9ucy5hdHRycyA6IHt9O1xuXG5cdC8vIEZvcmNlIHNpbmdsZS10YWcgc29sdXRpb24gb24gSUU2LTksIHdoaWNoIGhhcyBhIGhhcmQgbGltaXQgb24gdGhlICMgb2YgPHN0eWxlPlxuXHQvLyB0YWdzIGl0IHdpbGwgYWxsb3cgb24gYSBwYWdlXG5cdGlmICghb3B0aW9ucy5zaW5nbGV0b24gJiYgdHlwZW9mIG9wdGlvbnMuc2luZ2xldG9uICE9PSBcImJvb2xlYW5cIikgb3B0aW9ucy5zaW5nbGV0b24gPSBpc09sZElFKCk7XG5cblx0Ly8gQnkgZGVmYXVsdCwgYWRkIDxzdHlsZT4gdGFncyB0byB0aGUgPGhlYWQ+IGVsZW1lbnRcblx0aWYgKCFvcHRpb25zLmluc2VydEludG8pIG9wdGlvbnMuaW5zZXJ0SW50byA9IFwiaGVhZFwiO1xuXG5cdC8vIEJ5IGRlZmF1bHQsIGFkZCA8c3R5bGU+IHRhZ3MgdG8gdGhlIGJvdHRvbSBvZiB0aGUgdGFyZ2V0XG5cdGlmICghb3B0aW9ucy5pbnNlcnRBdCkgb3B0aW9ucy5pbnNlcnRBdCA9IFwiYm90dG9tXCI7XG5cblx0dmFyIHN0eWxlcyA9IGxpc3RUb1N0eWxlcyhsaXN0LCBvcHRpb25zKTtcblxuXHRhZGRTdHlsZXNUb0RvbShzdHlsZXMsIG9wdGlvbnMpO1xuXG5cdHJldHVybiBmdW5jdGlvbiB1cGRhdGUgKG5ld0xpc3QpIHtcblx0XHR2YXIgbWF5UmVtb3ZlID0gW107XG5cblx0XHRmb3IgKHZhciBpID0gMDsgaSA8IHN0eWxlcy5sZW5ndGg7IGkrKykge1xuXHRcdFx0dmFyIGl0ZW0gPSBzdHlsZXNbaV07XG5cdFx0XHR2YXIgZG9tU3R5bGUgPSBzdHlsZXNJbkRvbVtpdGVtLmlkXTtcblxuXHRcdFx0ZG9tU3R5bGUucmVmcy0tO1xuXHRcdFx0bWF5UmVtb3ZlLnB1c2goZG9tU3R5bGUpO1xuXHRcdH1cblxuXHRcdGlmKG5ld0xpc3QpIHtcblx0XHRcdHZhciBuZXdTdHlsZXMgPSBsaXN0VG9TdHlsZXMobmV3TGlzdCwgb3B0aW9ucyk7XG5cdFx0XHRhZGRTdHlsZXNUb0RvbShuZXdTdHlsZXMsIG9wdGlvbnMpO1xuXHRcdH1cblxuXHRcdGZvciAodmFyIGkgPSAwOyBpIDwgbWF5UmVtb3ZlLmxlbmd0aDsgaSsrKSB7XG5cdFx0XHR2YXIgZG9tU3R5bGUgPSBtYXlSZW1vdmVbaV07XG5cblx0XHRcdGlmKGRvbVN0eWxlLnJlZnMgPT09IDApIHtcblx0XHRcdFx0Zm9yICh2YXIgaiA9IDA7IGogPCBkb21TdHlsZS5wYXJ0cy5sZW5ndGg7IGorKykgZG9tU3R5bGUucGFydHNbal0oKTtcblxuXHRcdFx0XHRkZWxldGUgc3R5bGVzSW5Eb21bZG9tU3R5bGUuaWRdO1xuXHRcdFx0fVxuXHRcdH1cblx0fTtcbn07XG5cbmZ1bmN0aW9uIGFkZFN0eWxlc1RvRG9tIChzdHlsZXMsIG9wdGlvbnMpIHtcblx0Zm9yICh2YXIgaSA9IDA7IGkgPCBzdHlsZXMubGVuZ3RoOyBpKyspIHtcblx0XHR2YXIgaXRlbSA9IHN0eWxlc1tpXTtcblx0XHR2YXIgZG9tU3R5bGUgPSBzdHlsZXNJbkRvbVtpdGVtLmlkXTtcblxuXHRcdGlmKGRvbVN0eWxlKSB7XG5cdFx0XHRkb21TdHlsZS5yZWZzKys7XG5cblx0XHRcdGZvcih2YXIgaiA9IDA7IGogPCBkb21TdHlsZS5wYXJ0cy5sZW5ndGg7IGorKykge1xuXHRcdFx0XHRkb21TdHlsZS5wYXJ0c1tqXShpdGVtLnBhcnRzW2pdKTtcblx0XHRcdH1cblxuXHRcdFx0Zm9yKDsgaiA8IGl0ZW0ucGFydHMubGVuZ3RoOyBqKyspIHtcblx0XHRcdFx0ZG9tU3R5bGUucGFydHMucHVzaChhZGRTdHlsZShpdGVtLnBhcnRzW2pdLCBvcHRpb25zKSk7XG5cdFx0XHR9XG5cdFx0fSBlbHNlIHtcblx0XHRcdHZhciBwYXJ0cyA9IFtdO1xuXG5cdFx0XHRmb3IodmFyIGogPSAwOyBqIDwgaXRlbS5wYXJ0cy5sZW5ndGg7IGorKykge1xuXHRcdFx0XHRwYXJ0cy5wdXNoKGFkZFN0eWxlKGl0ZW0ucGFydHNbal0sIG9wdGlvbnMpKTtcblx0XHRcdH1cblxuXHRcdFx0c3R5bGVzSW5Eb21baXRlbS5pZF0gPSB7aWQ6IGl0ZW0uaWQsIHJlZnM6IDEsIHBhcnRzOiBwYXJ0c307XG5cdFx0fVxuXHR9XG59XG5cbmZ1bmN0aW9uIGxpc3RUb1N0eWxlcyAobGlzdCwgb3B0aW9ucykge1xuXHR2YXIgc3R5bGVzID0gW107XG5cdHZhciBuZXdTdHlsZXMgPSB7fTtcblxuXHRmb3IgKHZhciBpID0gMDsgaSA8IGxpc3QubGVuZ3RoOyBpKyspIHtcblx0XHR2YXIgaXRlbSA9IGxpc3RbaV07XG5cdFx0dmFyIGlkID0gb3B0aW9ucy5iYXNlID8gaXRlbVswXSArIG9wdGlvbnMuYmFzZSA6IGl0ZW1bMF07XG5cdFx0dmFyIGNzcyA9IGl0ZW1bMV07XG5cdFx0dmFyIG1lZGlhID0gaXRlbVsyXTtcblx0XHR2YXIgc291cmNlTWFwID0gaXRlbVszXTtcblx0XHR2YXIgcGFydCA9IHtjc3M6IGNzcywgbWVkaWE6IG1lZGlhLCBzb3VyY2VNYXA6IHNvdXJjZU1hcH07XG5cblx0XHRpZighbmV3U3R5bGVzW2lkXSkgc3R5bGVzLnB1c2gobmV3U3R5bGVzW2lkXSA9IHtpZDogaWQsIHBhcnRzOiBbcGFydF19KTtcblx0XHRlbHNlIG5ld1N0eWxlc1tpZF0ucGFydHMucHVzaChwYXJ0KTtcblx0fVxuXG5cdHJldHVybiBzdHlsZXM7XG59XG5cbmZ1bmN0aW9uIGluc2VydFN0eWxlRWxlbWVudCAob3B0aW9ucywgc3R5bGUpIHtcblx0dmFyIHRhcmdldCA9IGdldEVsZW1lbnQob3B0aW9ucy5pbnNlcnRJbnRvKVxuXG5cdGlmICghdGFyZ2V0KSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiQ291bGRuJ3QgZmluZCBhIHN0eWxlIHRhcmdldC4gVGhpcyBwcm9iYWJseSBtZWFucyB0aGF0IHRoZSB2YWx1ZSBmb3IgdGhlICdpbnNlcnRJbnRvJyBwYXJhbWV0ZXIgaXMgaW52YWxpZC5cIik7XG5cdH1cblxuXHR2YXIgbGFzdFN0eWxlRWxlbWVudEluc2VydGVkQXRUb3AgPSBzdHlsZXNJbnNlcnRlZEF0VG9wW3N0eWxlc0luc2VydGVkQXRUb3AubGVuZ3RoIC0gMV07XG5cblx0aWYgKG9wdGlvbnMuaW5zZXJ0QXQgPT09IFwidG9wXCIpIHtcblx0XHRpZiAoIWxhc3RTdHlsZUVsZW1lbnRJbnNlcnRlZEF0VG9wKSB7XG5cdFx0XHR0YXJnZXQuaW5zZXJ0QmVmb3JlKHN0eWxlLCB0YXJnZXQuZmlyc3RDaGlsZCk7XG5cdFx0fSBlbHNlIGlmIChsYXN0U3R5bGVFbGVtZW50SW5zZXJ0ZWRBdFRvcC5uZXh0U2libGluZykge1xuXHRcdFx0dGFyZ2V0Lmluc2VydEJlZm9yZShzdHlsZSwgbGFzdFN0eWxlRWxlbWVudEluc2VydGVkQXRUb3AubmV4dFNpYmxpbmcpO1xuXHRcdH0gZWxzZSB7XG5cdFx0XHR0YXJnZXQuYXBwZW5kQ2hpbGQoc3R5bGUpO1xuXHRcdH1cblx0XHRzdHlsZXNJbnNlcnRlZEF0VG9wLnB1c2goc3R5bGUpO1xuXHR9IGVsc2UgaWYgKG9wdGlvbnMuaW5zZXJ0QXQgPT09IFwiYm90dG9tXCIpIHtcblx0XHR0YXJnZXQuYXBwZW5kQ2hpbGQoc3R5bGUpO1xuXHR9IGVsc2UgaWYgKHR5cGVvZiBvcHRpb25zLmluc2VydEF0ID09PSBcIm9iamVjdFwiICYmIG9wdGlvbnMuaW5zZXJ0QXQuYmVmb3JlKSB7XG5cdFx0dmFyIG5leHRTaWJsaW5nID0gZ2V0RWxlbWVudChvcHRpb25zLmluc2VydEludG8gKyBcIiBcIiArIG9wdGlvbnMuaW5zZXJ0QXQuYmVmb3JlKTtcblx0XHR0YXJnZXQuaW5zZXJ0QmVmb3JlKHN0eWxlLCBuZXh0U2libGluZyk7XG5cdH0gZWxzZSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiW1N0eWxlIExvYWRlcl1cXG5cXG4gSW52YWxpZCB2YWx1ZSBmb3IgcGFyYW1ldGVyICdpbnNlcnRBdCcgKCdvcHRpb25zLmluc2VydEF0JykgZm91bmQuXFxuIE11c3QgYmUgJ3RvcCcsICdib3R0b20nLCBvciBPYmplY3QuXFxuIChodHRwczovL2dpdGh1Yi5jb20vd2VicGFjay1jb250cmliL3N0eWxlLWxvYWRlciNpbnNlcnRhdClcXG5cIik7XG5cdH1cbn1cblxuZnVuY3Rpb24gcmVtb3ZlU3R5bGVFbGVtZW50IChzdHlsZSkge1xuXHRpZiAoc3R5bGUucGFyZW50Tm9kZSA9PT0gbnVsbCkgcmV0dXJuIGZhbHNlO1xuXHRzdHlsZS5wYXJlbnROb2RlLnJlbW92ZUNoaWxkKHN0eWxlKTtcblxuXHR2YXIgaWR4ID0gc3R5bGVzSW5zZXJ0ZWRBdFRvcC5pbmRleE9mKHN0eWxlKTtcblx0aWYoaWR4ID49IDApIHtcblx0XHRzdHlsZXNJbnNlcnRlZEF0VG9wLnNwbGljZShpZHgsIDEpO1xuXHR9XG59XG5cbmZ1bmN0aW9uIGNyZWF0ZVN0eWxlRWxlbWVudCAob3B0aW9ucykge1xuXHR2YXIgc3R5bGUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3R5bGVcIik7XG5cblx0b3B0aW9ucy5hdHRycy50eXBlID0gXCJ0ZXh0L2Nzc1wiO1xuXG5cdGFkZEF0dHJzKHN0eWxlLCBvcHRpb25zLmF0dHJzKTtcblx0aW5zZXJ0U3R5bGVFbGVtZW50KG9wdGlvbnMsIHN0eWxlKTtcblxuXHRyZXR1cm4gc3R5bGU7XG59XG5cbmZ1bmN0aW9uIGNyZWF0ZUxpbmtFbGVtZW50IChvcHRpb25zKSB7XG5cdHZhciBsaW5rID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImxpbmtcIik7XG5cblx0b3B0aW9ucy5hdHRycy50eXBlID0gXCJ0ZXh0L2Nzc1wiO1xuXHRvcHRpb25zLmF0dHJzLnJlbCA9IFwic3R5bGVzaGVldFwiO1xuXG5cdGFkZEF0dHJzKGxpbmssIG9wdGlvbnMuYXR0cnMpO1xuXHRpbnNlcnRTdHlsZUVsZW1lbnQob3B0aW9ucywgbGluayk7XG5cblx0cmV0dXJuIGxpbms7XG59XG5cbmZ1bmN0aW9uIGFkZEF0dHJzIChlbCwgYXR0cnMpIHtcblx0T2JqZWN0LmtleXMoYXR0cnMpLmZvckVhY2goZnVuY3Rpb24gKGtleSkge1xuXHRcdGVsLnNldEF0dHJpYnV0ZShrZXksIGF0dHJzW2tleV0pO1xuXHR9KTtcbn1cblxuZnVuY3Rpb24gYWRkU3R5bGUgKG9iaiwgb3B0aW9ucykge1xuXHR2YXIgc3R5bGUsIHVwZGF0ZSwgcmVtb3ZlLCByZXN1bHQ7XG5cblx0Ly8gSWYgYSB0cmFuc2Zvcm0gZnVuY3Rpb24gd2FzIGRlZmluZWQsIHJ1biBpdCBvbiB0aGUgY3NzXG5cdGlmIChvcHRpb25zLnRyYW5zZm9ybSAmJiBvYmouY3NzKSB7XG5cdCAgICByZXN1bHQgPSBvcHRpb25zLnRyYW5zZm9ybShvYmouY3NzKTtcblxuXHQgICAgaWYgKHJlc3VsdCkge1xuXHQgICAgXHQvLyBJZiB0cmFuc2Zvcm0gcmV0dXJucyBhIHZhbHVlLCB1c2UgdGhhdCBpbnN0ZWFkIG9mIHRoZSBvcmlnaW5hbCBjc3MuXG5cdCAgICBcdC8vIFRoaXMgYWxsb3dzIHJ1bm5pbmcgcnVudGltZSB0cmFuc2Zvcm1hdGlvbnMgb24gdGhlIGNzcy5cblx0ICAgIFx0b2JqLmNzcyA9IHJlc3VsdDtcblx0ICAgIH0gZWxzZSB7XG5cdCAgICBcdC8vIElmIHRoZSB0cmFuc2Zvcm0gZnVuY3Rpb24gcmV0dXJucyBhIGZhbHN5IHZhbHVlLCBkb24ndCBhZGQgdGhpcyBjc3MuXG5cdCAgICBcdC8vIFRoaXMgYWxsb3dzIGNvbmRpdGlvbmFsIGxvYWRpbmcgb2YgY3NzXG5cdCAgICBcdHJldHVybiBmdW5jdGlvbigpIHtcblx0ICAgIFx0XHQvLyBub29wXG5cdCAgICBcdH07XG5cdCAgICB9XG5cdH1cblxuXHRpZiAob3B0aW9ucy5zaW5nbGV0b24pIHtcblx0XHR2YXIgc3R5bGVJbmRleCA9IHNpbmdsZXRvbkNvdW50ZXIrKztcblxuXHRcdHN0eWxlID0gc2luZ2xldG9uIHx8IChzaW5nbGV0b24gPSBjcmVhdGVTdHlsZUVsZW1lbnQob3B0aW9ucykpO1xuXG5cdFx0dXBkYXRlID0gYXBwbHlUb1NpbmdsZXRvblRhZy5iaW5kKG51bGwsIHN0eWxlLCBzdHlsZUluZGV4LCBmYWxzZSk7XG5cdFx0cmVtb3ZlID0gYXBwbHlUb1NpbmdsZXRvblRhZy5iaW5kKG51bGwsIHN0eWxlLCBzdHlsZUluZGV4LCB0cnVlKTtcblxuXHR9IGVsc2UgaWYgKFxuXHRcdG9iai5zb3VyY2VNYXAgJiZcblx0XHR0eXBlb2YgVVJMID09PSBcImZ1bmN0aW9uXCIgJiZcblx0XHR0eXBlb2YgVVJMLmNyZWF0ZU9iamVjdFVSTCA9PT0gXCJmdW5jdGlvblwiICYmXG5cdFx0dHlwZW9mIFVSTC5yZXZva2VPYmplY3RVUkwgPT09IFwiZnVuY3Rpb25cIiAmJlxuXHRcdHR5cGVvZiBCbG9iID09PSBcImZ1bmN0aW9uXCIgJiZcblx0XHR0eXBlb2YgYnRvYSA9PT0gXCJmdW5jdGlvblwiXG5cdCkge1xuXHRcdHN0eWxlID0gY3JlYXRlTGlua0VsZW1lbnQob3B0aW9ucyk7XG5cdFx0dXBkYXRlID0gdXBkYXRlTGluay5iaW5kKG51bGwsIHN0eWxlLCBvcHRpb25zKTtcblx0XHRyZW1vdmUgPSBmdW5jdGlvbiAoKSB7XG5cdFx0XHRyZW1vdmVTdHlsZUVsZW1lbnQoc3R5bGUpO1xuXG5cdFx0XHRpZihzdHlsZS5ocmVmKSBVUkwucmV2b2tlT2JqZWN0VVJMKHN0eWxlLmhyZWYpO1xuXHRcdH07XG5cdH0gZWxzZSB7XG5cdFx0c3R5bGUgPSBjcmVhdGVTdHlsZUVsZW1lbnQob3B0aW9ucyk7XG5cdFx0dXBkYXRlID0gYXBwbHlUb1RhZy5iaW5kKG51bGwsIHN0eWxlKTtcblx0XHRyZW1vdmUgPSBmdW5jdGlvbiAoKSB7XG5cdFx0XHRyZW1vdmVTdHlsZUVsZW1lbnQoc3R5bGUpO1xuXHRcdH07XG5cdH1cblxuXHR1cGRhdGUob2JqKTtcblxuXHRyZXR1cm4gZnVuY3Rpb24gdXBkYXRlU3R5bGUgKG5ld09iaikge1xuXHRcdGlmIChuZXdPYmopIHtcblx0XHRcdGlmIChcblx0XHRcdFx0bmV3T2JqLmNzcyA9PT0gb2JqLmNzcyAmJlxuXHRcdFx0XHRuZXdPYmoubWVkaWEgPT09IG9iai5tZWRpYSAmJlxuXHRcdFx0XHRuZXdPYmouc291cmNlTWFwID09PSBvYmouc291cmNlTWFwXG5cdFx0XHQpIHtcblx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0fVxuXG5cdFx0XHR1cGRhdGUob2JqID0gbmV3T2JqKTtcblx0XHR9IGVsc2Uge1xuXHRcdFx0cmVtb3ZlKCk7XG5cdFx0fVxuXHR9O1xufVxuXG52YXIgcmVwbGFjZVRleHQgPSAoZnVuY3Rpb24gKCkge1xuXHR2YXIgdGV4dFN0b3JlID0gW107XG5cblx0cmV0dXJuIGZ1bmN0aW9uIChpbmRleCwgcmVwbGFjZW1lbnQpIHtcblx0XHR0ZXh0U3RvcmVbaW5kZXhdID0gcmVwbGFjZW1lbnQ7XG5cblx0XHRyZXR1cm4gdGV4dFN0b3JlLmZpbHRlcihCb29sZWFuKS5qb2luKCdcXG4nKTtcblx0fTtcbn0pKCk7XG5cbmZ1bmN0aW9uIGFwcGx5VG9TaW5nbGV0b25UYWcgKHN0eWxlLCBpbmRleCwgcmVtb3ZlLCBvYmopIHtcblx0dmFyIGNzcyA9IHJlbW92ZSA/IFwiXCIgOiBvYmouY3NzO1xuXG5cdGlmIChzdHlsZS5zdHlsZVNoZWV0KSB7XG5cdFx0c3R5bGUuc3R5bGVTaGVldC5jc3NUZXh0ID0gcmVwbGFjZVRleHQoaW5kZXgsIGNzcyk7XG5cdH0gZWxzZSB7XG5cdFx0dmFyIGNzc05vZGUgPSBkb2N1bWVudC5jcmVhdGVUZXh0Tm9kZShjc3MpO1xuXHRcdHZhciBjaGlsZE5vZGVzID0gc3R5bGUuY2hpbGROb2RlcztcblxuXHRcdGlmIChjaGlsZE5vZGVzW2luZGV4XSkgc3R5bGUucmVtb3ZlQ2hpbGQoY2hpbGROb2Rlc1tpbmRleF0pO1xuXG5cdFx0aWYgKGNoaWxkTm9kZXMubGVuZ3RoKSB7XG5cdFx0XHRzdHlsZS5pbnNlcnRCZWZvcmUoY3NzTm9kZSwgY2hpbGROb2Rlc1tpbmRleF0pO1xuXHRcdH0gZWxzZSB7XG5cdFx0XHRzdHlsZS5hcHBlbmRDaGlsZChjc3NOb2RlKTtcblx0XHR9XG5cdH1cbn1cblxuZnVuY3Rpb24gYXBwbHlUb1RhZyAoc3R5bGUsIG9iaikge1xuXHR2YXIgY3NzID0gb2JqLmNzcztcblx0dmFyIG1lZGlhID0gb2JqLm1lZGlhO1xuXG5cdGlmKG1lZGlhKSB7XG5cdFx0c3R5bGUuc2V0QXR0cmlidXRlKFwibWVkaWFcIiwgbWVkaWEpXG5cdH1cblxuXHRpZihzdHlsZS5zdHlsZVNoZWV0KSB7XG5cdFx0c3R5bGUuc3R5bGVTaGVldC5jc3NUZXh0ID0gY3NzO1xuXHR9IGVsc2Uge1xuXHRcdHdoaWxlKHN0eWxlLmZpcnN0Q2hpbGQpIHtcblx0XHRcdHN0eWxlLnJlbW92ZUNoaWxkKHN0eWxlLmZpcnN0Q2hpbGQpO1xuXHRcdH1cblxuXHRcdHN0eWxlLmFwcGVuZENoaWxkKGRvY3VtZW50LmNyZWF0ZVRleHROb2RlKGNzcykpO1xuXHR9XG59XG5cbmZ1bmN0aW9uIHVwZGF0ZUxpbmsgKGxpbmssIG9wdGlvbnMsIG9iaikge1xuXHR2YXIgY3NzID0gb2JqLmNzcztcblx0dmFyIHNvdXJjZU1hcCA9IG9iai5zb3VyY2VNYXA7XG5cblx0Lypcblx0XHRJZiBjb252ZXJ0VG9BYnNvbHV0ZVVybHMgaXNuJ3QgZGVmaW5lZCwgYnV0IHNvdXJjZW1hcHMgYXJlIGVuYWJsZWRcblx0XHRhbmQgdGhlcmUgaXMgbm8gcHVibGljUGF0aCBkZWZpbmVkIHRoZW4gbGV0cyB0dXJuIGNvbnZlcnRUb0Fic29sdXRlVXJsc1xuXHRcdG9uIGJ5IGRlZmF1bHQuICBPdGhlcndpc2UgZGVmYXVsdCB0byB0aGUgY29udmVydFRvQWJzb2x1dGVVcmxzIG9wdGlvblxuXHRcdGRpcmVjdGx5XG5cdCovXG5cdHZhciBhdXRvRml4VXJscyA9IG9wdGlvbnMuY29udmVydFRvQWJzb2x1dGVVcmxzID09PSB1bmRlZmluZWQgJiYgc291cmNlTWFwO1xuXG5cdGlmIChvcHRpb25zLmNvbnZlcnRUb0Fic29sdXRlVXJscyB8fCBhdXRvRml4VXJscykge1xuXHRcdGNzcyA9IGZpeFVybHMoY3NzKTtcblx0fVxuXG5cdGlmIChzb3VyY2VNYXApIHtcblx0XHQvLyBodHRwOi8vc3RhY2tvdmVyZmxvdy5jb20vYS8yNjYwMzg3NVxuXHRcdGNzcyArPSBcIlxcbi8qIyBzb3VyY2VNYXBwaW5nVVJMPWRhdGE6YXBwbGljYXRpb24vanNvbjtiYXNlNjQsXCIgKyBidG9hKHVuZXNjYXBlKGVuY29kZVVSSUNvbXBvbmVudChKU09OLnN0cmluZ2lmeShzb3VyY2VNYXApKSkpICsgXCIgKi9cIjtcblx0fVxuXG5cdHZhciBibG9iID0gbmV3IEJsb2IoW2Nzc10sIHsgdHlwZTogXCJ0ZXh0L2Nzc1wiIH0pO1xuXG5cdHZhciBvbGRTcmMgPSBsaW5rLmhyZWY7XG5cblx0bGluay5ocmVmID0gVVJMLmNyZWF0ZU9iamVjdFVSTChibG9iKTtcblxuXHRpZihvbGRTcmMpIFVSTC5yZXZva2VPYmplY3RVUkwob2xkU3JjKTtcbn1cblxuXG5cbi8vLy8vLy8vLy8vLy8vLy8vL1xuLy8gV0VCUEFDSyBGT09URVJcbi8vIC4vbm9kZV9tb2R1bGVzL3N0eWxlLWxvYWRlci9saWIvYWRkU3R5bGVzLmpzXG4vLyBtb2R1bGUgaWQgPSA2XG4vLyBtb2R1bGUgY2h1bmtzID0gMCIsIlxuLyoqXG4gKiBXaGVuIHNvdXJjZSBtYXBzIGFyZSBlbmFibGVkLCBgc3R5bGUtbG9hZGVyYCB1c2VzIGEgbGluayBlbGVtZW50IHdpdGggYSBkYXRhLXVyaSB0b1xuICogZW1iZWQgdGhlIGNzcyBvbiB0aGUgcGFnZS4gVGhpcyBicmVha3MgYWxsIHJlbGF0aXZlIHVybHMgYmVjYXVzZSBub3cgdGhleSBhcmUgcmVsYXRpdmUgdG8gYVxuICogYnVuZGxlIGluc3RlYWQgb2YgdGhlIGN1cnJlbnQgcGFnZS5cbiAqXG4gKiBPbmUgc29sdXRpb24gaXMgdG8gb25seSB1c2UgZnVsbCB1cmxzLCBidXQgdGhhdCBtYXkgYmUgaW1wb3NzaWJsZS5cbiAqXG4gKiBJbnN0ZWFkLCB0aGlzIGZ1bmN0aW9uIFwiZml4ZXNcIiB0aGUgcmVsYXRpdmUgdXJscyB0byBiZSBhYnNvbHV0ZSBhY2NvcmRpbmcgdG8gdGhlIGN1cnJlbnQgcGFnZSBsb2NhdGlvbi5cbiAqXG4gKiBBIHJ1ZGltZW50YXJ5IHRlc3Qgc3VpdGUgaXMgbG9jYXRlZCBhdCBgdGVzdC9maXhVcmxzLmpzYCBhbmQgY2FuIGJlIHJ1biB2aWEgdGhlIGBucG0gdGVzdGAgY29tbWFuZC5cbiAqXG4gKi9cblxubW9kdWxlLmV4cG9ydHMgPSBmdW5jdGlvbiAoY3NzKSB7XG4gIC8vIGdldCBjdXJyZW50IGxvY2F0aW9uXG4gIHZhciBsb2NhdGlvbiA9IHR5cGVvZiB3aW5kb3cgIT09IFwidW5kZWZpbmVkXCIgJiYgd2luZG93LmxvY2F0aW9uO1xuXG4gIGlmICghbG9jYXRpb24pIHtcbiAgICB0aHJvdyBuZXcgRXJyb3IoXCJmaXhVcmxzIHJlcXVpcmVzIHdpbmRvdy5sb2NhdGlvblwiKTtcbiAgfVxuXG5cdC8vIGJsYW5rIG9yIG51bGw/XG5cdGlmICghY3NzIHx8IHR5cGVvZiBjc3MgIT09IFwic3RyaW5nXCIpIHtcblx0ICByZXR1cm4gY3NzO1xuICB9XG5cbiAgdmFyIGJhc2VVcmwgPSBsb2NhdGlvbi5wcm90b2NvbCArIFwiLy9cIiArIGxvY2F0aW9uLmhvc3Q7XG4gIHZhciBjdXJyZW50RGlyID0gYmFzZVVybCArIGxvY2F0aW9uLnBhdGhuYW1lLnJlcGxhY2UoL1xcL1teXFwvXSokLywgXCIvXCIpO1xuXG5cdC8vIGNvbnZlcnQgZWFjaCB1cmwoLi4uKVxuXHQvKlxuXHRUaGlzIHJlZ3VsYXIgZXhwcmVzc2lvbiBpcyBqdXN0IGEgd2F5IHRvIHJlY3Vyc2l2ZWx5IG1hdGNoIGJyYWNrZXRzIHdpdGhpblxuXHRhIHN0cmluZy5cblxuXHQgL3VybFxccypcXCggID0gTWF0Y2ggb24gdGhlIHdvcmQgXCJ1cmxcIiB3aXRoIGFueSB3aGl0ZXNwYWNlIGFmdGVyIGl0IGFuZCB0aGVuIGEgcGFyZW5zXG5cdCAgICggID0gU3RhcnQgYSBjYXB0dXJpbmcgZ3JvdXBcblx0ICAgICAoPzogID0gU3RhcnQgYSBub24tY2FwdHVyaW5nIGdyb3VwXG5cdCAgICAgICAgIFteKShdICA9IE1hdGNoIGFueXRoaW5nIHRoYXQgaXNuJ3QgYSBwYXJlbnRoZXNlc1xuXHQgICAgICAgICB8ICA9IE9SXG5cdCAgICAgICAgIFxcKCAgPSBNYXRjaCBhIHN0YXJ0IHBhcmVudGhlc2VzXG5cdCAgICAgICAgICAgICAoPzogID0gU3RhcnQgYW5vdGhlciBub24tY2FwdHVyaW5nIGdyb3Vwc1xuXHQgICAgICAgICAgICAgICAgIFteKShdKyAgPSBNYXRjaCBhbnl0aGluZyB0aGF0IGlzbid0IGEgcGFyZW50aGVzZXNcblx0ICAgICAgICAgICAgICAgICB8ICA9IE9SXG5cdCAgICAgICAgICAgICAgICAgXFwoICA9IE1hdGNoIGEgc3RhcnQgcGFyZW50aGVzZXNcblx0ICAgICAgICAgICAgICAgICAgICAgW14pKF0qICA9IE1hdGNoIGFueXRoaW5nIHRoYXQgaXNuJ3QgYSBwYXJlbnRoZXNlc1xuXHQgICAgICAgICAgICAgICAgIFxcKSAgPSBNYXRjaCBhIGVuZCBwYXJlbnRoZXNlc1xuXHQgICAgICAgICAgICAgKSAgPSBFbmQgR3JvdXBcbiAgICAgICAgICAgICAgKlxcKSA9IE1hdGNoIGFueXRoaW5nIGFuZCB0aGVuIGEgY2xvc2UgcGFyZW5zXG4gICAgICAgICAgKSAgPSBDbG9zZSBub24tY2FwdHVyaW5nIGdyb3VwXG4gICAgICAgICAgKiAgPSBNYXRjaCBhbnl0aGluZ1xuICAgICAgICkgID0gQ2xvc2UgY2FwdHVyaW5nIGdyb3VwXG5cdCBcXCkgID0gTWF0Y2ggYSBjbG9zZSBwYXJlbnNcblxuXHQgL2dpICA9IEdldCBhbGwgbWF0Y2hlcywgbm90IHRoZSBmaXJzdC4gIEJlIGNhc2UgaW5zZW5zaXRpdmUuXG5cdCAqL1xuXHR2YXIgZml4ZWRDc3MgPSBjc3MucmVwbGFjZSgvdXJsXFxzKlxcKCgoPzpbXikoXXxcXCgoPzpbXikoXSt8XFwoW14pKF0qXFwpKSpcXCkpKilcXCkvZ2ksIGZ1bmN0aW9uKGZ1bGxNYXRjaCwgb3JpZ1VybCkge1xuXHRcdC8vIHN0cmlwIHF1b3RlcyAoaWYgdGhleSBleGlzdClcblx0XHR2YXIgdW5xdW90ZWRPcmlnVXJsID0gb3JpZ1VybFxuXHRcdFx0LnRyaW0oKVxuXHRcdFx0LnJlcGxhY2UoL15cIiguKilcIiQvLCBmdW5jdGlvbihvLCAkMSl7IHJldHVybiAkMTsgfSlcblx0XHRcdC5yZXBsYWNlKC9eJyguKiknJC8sIGZ1bmN0aW9uKG8sICQxKXsgcmV0dXJuICQxOyB9KTtcblxuXHRcdC8vIGFscmVhZHkgYSBmdWxsIHVybD8gbm8gY2hhbmdlXG5cdFx0aWYgKC9eKCN8ZGF0YTp8aHR0cDpcXC9cXC98aHR0cHM6XFwvXFwvfGZpbGU6XFwvXFwvXFwvKS9pLnRlc3QodW5xdW90ZWRPcmlnVXJsKSkge1xuXHRcdCAgcmV0dXJuIGZ1bGxNYXRjaDtcblx0XHR9XG5cblx0XHQvLyBjb252ZXJ0IHRoZSB1cmwgdG8gYSBmdWxsIHVybFxuXHRcdHZhciBuZXdVcmw7XG5cblx0XHRpZiAodW5xdW90ZWRPcmlnVXJsLmluZGV4T2YoXCIvL1wiKSA9PT0gMCkge1xuXHRcdCAgXHQvL1RPRE86IHNob3VsZCB3ZSBhZGQgcHJvdG9jb2w/XG5cdFx0XHRuZXdVcmwgPSB1bnF1b3RlZE9yaWdVcmw7XG5cdFx0fSBlbHNlIGlmICh1bnF1b3RlZE9yaWdVcmwuaW5kZXhPZihcIi9cIikgPT09IDApIHtcblx0XHRcdC8vIHBhdGggc2hvdWxkIGJlIHJlbGF0aXZlIHRvIHRoZSBiYXNlIHVybFxuXHRcdFx0bmV3VXJsID0gYmFzZVVybCArIHVucXVvdGVkT3JpZ1VybDsgLy8gYWxyZWFkeSBzdGFydHMgd2l0aCAnLydcblx0XHR9IGVsc2Uge1xuXHRcdFx0Ly8gcGF0aCBzaG91bGQgYmUgcmVsYXRpdmUgdG8gY3VycmVudCBkaXJlY3Rvcnlcblx0XHRcdG5ld1VybCA9IGN1cnJlbnREaXIgKyB1bnF1b3RlZE9yaWdVcmwucmVwbGFjZSgvXlxcLlxcLy8sIFwiXCIpOyAvLyBTdHJpcCBsZWFkaW5nICcuLydcblx0XHR9XG5cblx0XHQvLyBzZW5kIGJhY2sgdGhlIGZpeGVkIHVybCguLi4pXG5cdFx0cmV0dXJuIFwidXJsKFwiICsgSlNPTi5zdHJpbmdpZnkobmV3VXJsKSArIFwiKVwiO1xuXHR9KTtcblxuXHQvLyBzZW5kIGJhY2sgdGhlIGZpeGVkIGNzc1xuXHRyZXR1cm4gZml4ZWRDc3M7XG59O1xuXG5cblxuLy8vLy8vLy8vLy8vLy8vLy8vXG4vLyBXRUJQQUNLIEZPT1RFUlxuLy8gLi9ub2RlX21vZHVsZXMvc3R5bGUtbG9hZGVyL2xpYi91cmxzLmpzXG4vLyBtb2R1bGUgaWQgPSA3XG4vLyBtb2R1bGUgY2h1bmtzID0gMCIsIi8vIENvcHlyaWdodCAyMDE4IFRoZSBHbyBBdXRob3JzLiBBbGwgcmlnaHRzIHJlc2VydmVkLlxuLy8gVXNlIG9mIHRoaXMgc291cmNlIGNvZGUgaXMgZ292ZXJuZWQgYnkgYSBCU0Qtc3R5bGVcbi8vIGxpY2Vuc2UgdGhhdCBjYW4gYmUgZm91bmQgaW4gdGhlIExJQ0VOU0UgZmlsZS5cblxuXCJ1c2Ugc3RyaWN0XCI7XG5cbigoKSA9PiB7XG5cdGNvbnN0IGVub3N5cyA9ICgpID0+IHtcblx0XHRjb25zdCBlcnIgPSBuZXcgRXJyb3IoXCJub3QgaW1wbGVtZW50ZWRcIik7XG5cdFx0ZXJyLmNvZGUgPSBcIkVOT1NZU1wiO1xuXHRcdHJldHVybiBlcnI7XG5cdH07XG5cblx0aWYgKCFnbG9iYWxUaGlzLmZzKSB7XG5cdFx0bGV0IG91dHB1dEJ1ZiA9IFwiXCI7XG5cdFx0Z2xvYmFsVGhpcy5mcyA9IHtcblx0XHRcdGNvbnN0YW50czogeyBPX1dST05MWTogLTEsIE9fUkRXUjogLTEsIE9fQ1JFQVQ6IC0xLCBPX1RSVU5DOiAtMSwgT19BUFBFTkQ6IC0xLCBPX0VYQ0w6IC0xIH0sIC8vIHVudXNlZFxuXHRcdFx0d3JpdGVTeW5jKGZkLCBidWYpIHtcblx0XHRcdFx0b3V0cHV0QnVmICs9IGRlY29kZXIuZGVjb2RlKGJ1Zik7XG5cdFx0XHRcdGNvbnN0IG5sID0gb3V0cHV0QnVmLmxhc3RJbmRleE9mKFwiXFxuXCIpO1xuXHRcdFx0XHRpZiAobmwgIT0gLTEpIHtcblx0XHRcdFx0XHRjb25zb2xlLmxvZyhvdXRwdXRCdWYuc3Vic3RyaW5nKDAsIG5sKSk7XG5cdFx0XHRcdFx0b3V0cHV0QnVmID0gb3V0cHV0QnVmLnN1YnN0cmluZyhubCArIDEpO1xuXHRcdFx0XHR9XG5cdFx0XHRcdHJldHVybiBidWYubGVuZ3RoO1xuXHRcdFx0fSxcblx0XHRcdHdyaXRlKGZkLCBidWYsIG9mZnNldCwgbGVuZ3RoLCBwb3NpdGlvbiwgY2FsbGJhY2spIHtcblx0XHRcdFx0aWYgKG9mZnNldCAhPT0gMCB8fCBsZW5ndGggIT09IGJ1Zi5sZW5ndGggfHwgcG9zaXRpb24gIT09IG51bGwpIHtcblx0XHRcdFx0XHRjYWxsYmFjayhlbm9zeXMoKSk7XG5cdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHR9XG5cdFx0XHRcdGNvbnN0IG4gPSB0aGlzLndyaXRlU3luYyhmZCwgYnVmKTtcblx0XHRcdFx0Y2FsbGJhY2sobnVsbCwgbik7XG5cdFx0XHR9LFxuXHRcdFx0Y2htb2QocGF0aCwgbW9kZSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0Y2hvd24ocGF0aCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGNsb3NlKGZkLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRmY2htb2QoZmQsIG1vZGUsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGZjaG93bihmZCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGZzdGF0KGZkLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRmc3luYyhmZCwgY2FsbGJhY2spIHsgY2FsbGJhY2sobnVsbCk7IH0sXG5cdFx0XHRmdHJ1bmNhdGUoZmQsIGxlbmd0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0bGNob3duKHBhdGgsIHVpZCwgZ2lkLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRsaW5rKHBhdGgsIGxpbmssIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGxzdGF0KHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdG1rZGlyKHBhdGgsIHBlcm0sIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdG9wZW4ocGF0aCwgZmxhZ3MsIG1vZGUsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHJlYWQoZmQsIGJ1ZmZlciwgb2Zmc2V0LCBsZW5ndGgsIHBvc2l0aW9uLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRyZWFkZGlyKHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHJlYWRsaW5rKHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHJlbmFtZShmcm9tLCB0bywgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0cm1kaXIocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0c3RhdChwYXRoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRzeW1saW5rKHBhdGgsIGxpbmssIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHRydW5jYXRlKHBhdGgsIGxlbmd0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0dW5saW5rKHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHV0aW1lcyhwYXRoLCBhdGltZSwgbXRpbWUsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHR9O1xuXHR9XG5cblx0aWYgKCFnbG9iYWxUaGlzLnByb2Nlc3MpIHtcblx0XHRnbG9iYWxUaGlzLnByb2Nlc3MgPSB7XG5cdFx0XHRnZXR1aWQoKSB7IHJldHVybiAtMTsgfSxcblx0XHRcdGdldGdpZCgpIHsgcmV0dXJuIC0xOyB9LFxuXHRcdFx0Z2V0ZXVpZCgpIHsgcmV0dXJuIC0xOyB9LFxuXHRcdFx0Z2V0ZWdpZCgpIHsgcmV0dXJuIC0xOyB9LFxuXHRcdFx0Z2V0Z3JvdXBzKCkgeyB0aHJvdyBlbm9zeXMoKTsgfSxcblx0XHRcdHBpZDogLTEsXG5cdFx0XHRwcGlkOiAtMSxcblx0XHRcdHVtYXNrKCkgeyB0aHJvdyBlbm9zeXMoKTsgfSxcblx0XHRcdGN3ZCgpIHsgdGhyb3cgZW5vc3lzKCk7IH0sXG5cdFx0XHRjaGRpcigpIHsgdGhyb3cgZW5vc3lzKCk7IH0sXG5cdFx0fVxuXHR9XG5cblx0aWYgKCFnbG9iYWxUaGlzLmNyeXB0bykge1xuXHRcdHRocm93IG5ldyBFcnJvcihcImdsb2JhbFRoaXMuY3J5cHRvIGlzIG5vdCBhdmFpbGFibGUsIHBvbHlmaWxsIHJlcXVpcmVkIChjcnlwdG8uZ2V0UmFuZG9tVmFsdWVzIG9ubHkpXCIpO1xuXHR9XG5cblx0aWYgKCFnbG9iYWxUaGlzLnBlcmZvcm1hbmNlKSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiZ2xvYmFsVGhpcy5wZXJmb3JtYW5jZSBpcyBub3QgYXZhaWxhYmxlLCBwb2x5ZmlsbCByZXF1aXJlZCAocGVyZm9ybWFuY2Uubm93IG9ubHkpXCIpO1xuXHR9XG5cblx0aWYgKCFnbG9iYWxUaGlzLlRleHRFbmNvZGVyKSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiZ2xvYmFsVGhpcy5UZXh0RW5jb2RlciBpcyBub3QgYXZhaWxhYmxlLCBwb2x5ZmlsbCByZXF1aXJlZFwiKTtcblx0fVxuXG5cdGlmICghZ2xvYmFsVGhpcy5UZXh0RGVjb2Rlcikge1xuXHRcdHRocm93IG5ldyBFcnJvcihcImdsb2JhbFRoaXMuVGV4dERlY29kZXIgaXMgbm90IGF2YWlsYWJsZSwgcG9seWZpbGwgcmVxdWlyZWRcIik7XG5cdH1cblxuXHRjb25zdCBlbmNvZGVyID0gbmV3IFRleHRFbmNvZGVyKFwidXRmLThcIik7XG5cdGNvbnN0IGRlY29kZXIgPSBuZXcgVGV4dERlY29kZXIoXCJ1dGYtOFwiKTtcblxuXHRnbG9iYWxUaGlzLkdvID0gY2xhc3Mge1xuXHRcdGNvbnN0cnVjdG9yKCkge1xuXHRcdFx0dGhpcy5hcmd2ID0gW1wianNcIl07XG5cdFx0XHR0aGlzLmVudiA9IHt9O1xuXHRcdFx0dGhpcy5leGl0ID0gKGNvZGUpID0+IHtcblx0XHRcdFx0aWYgKGNvZGUgIT09IDApIHtcblx0XHRcdFx0XHRjb25zb2xlLndhcm4oXCJleGl0IGNvZGU6XCIsIGNvZGUpO1xuXHRcdFx0XHR9XG5cdFx0XHR9O1xuXHRcdFx0dGhpcy5fZXhpdFByb21pc2UgPSBuZXcgUHJvbWlzZSgocmVzb2x2ZSkgPT4ge1xuXHRcdFx0XHR0aGlzLl9yZXNvbHZlRXhpdFByb21pc2UgPSByZXNvbHZlO1xuXHRcdFx0fSk7XG5cdFx0XHR0aGlzLl9wZW5kaW5nRXZlbnQgPSBudWxsO1xuXHRcdFx0dGhpcy5fc2NoZWR1bGVkVGltZW91dHMgPSBuZXcgTWFwKCk7XG5cdFx0XHR0aGlzLl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQgPSAxO1xuXG5cdFx0XHRjb25zdCBzZXRJbnQ2NCA9IChhZGRyLCB2KSA9PiB7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihhZGRyICsgMCwgdiwgdHJ1ZSk7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihhZGRyICsgNCwgTWF0aC5mbG9vcih2IC8gNDI5NDk2NzI5NiksIHRydWUpO1xuXHRcdFx0fVxuXG5cdFx0XHRjb25zdCBzZXRJbnQzMiA9IChhZGRyLCB2KSA9PiB7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihhZGRyICsgMCwgdiwgdHJ1ZSk7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IGdldEludDY0ID0gKGFkZHIpID0+IHtcblx0XHRcdFx0Y29uc3QgbG93ID0gdGhpcy5tZW0uZ2V0VWludDMyKGFkZHIgKyAwLCB0cnVlKTtcblx0XHRcdFx0Y29uc3QgaGlnaCA9IHRoaXMubWVtLmdldEludDMyKGFkZHIgKyA0LCB0cnVlKTtcblx0XHRcdFx0cmV0dXJuIGxvdyArIGhpZ2ggKiA0Mjk0OTY3Mjk2O1xuXHRcdFx0fVxuXG5cdFx0XHRjb25zdCBsb2FkVmFsdWUgPSAoYWRkcikgPT4ge1xuXHRcdFx0XHRjb25zdCBmID0gdGhpcy5tZW0uZ2V0RmxvYXQ2NChhZGRyLCB0cnVlKTtcblx0XHRcdFx0aWYgKGYgPT09IDApIHtcblx0XHRcdFx0XHRyZXR1cm4gdW5kZWZpbmVkO1xuXHRcdFx0XHR9XG5cdFx0XHRcdGlmICghaXNOYU4oZikpIHtcblx0XHRcdFx0XHRyZXR1cm4gZjtcblx0XHRcdFx0fVxuXG5cdFx0XHRcdGNvbnN0IGlkID0gdGhpcy5tZW0uZ2V0VWludDMyKGFkZHIsIHRydWUpO1xuXHRcdFx0XHRyZXR1cm4gdGhpcy5fdmFsdWVzW2lkXTtcblx0XHRcdH1cblxuXHRcdFx0Y29uc3Qgc3RvcmVWYWx1ZSA9IChhZGRyLCB2KSA9PiB7XG5cdFx0XHRcdGNvbnN0IG5hbkhlYWQgPSAweDdGRjgwMDAwO1xuXG5cdFx0XHRcdGlmICh0eXBlb2YgdiA9PT0gXCJudW1iZXJcIiAmJiB2ICE9PSAwKSB7XG5cdFx0XHRcdFx0aWYgKGlzTmFOKHYpKSB7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpO1xuXHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDMyKGFkZHIsIDAsIHRydWUpO1xuXHRcdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHRcdH1cblx0XHRcdFx0XHR0aGlzLm1lbS5zZXRGbG9hdDY0KGFkZHIsIHYsIHRydWUpO1xuXHRcdFx0XHRcdHJldHVybjtcblx0XHRcdFx0fVxuXG5cdFx0XHRcdGlmICh2ID09PSB1bmRlZmluZWQpIHtcblx0XHRcdFx0XHR0aGlzLm1lbS5zZXRGbG9hdDY0KGFkZHIsIDAsIHRydWUpO1xuXHRcdFx0XHRcdHJldHVybjtcblx0XHRcdFx0fVxuXG5cdFx0XHRcdGxldCBpZCA9IHRoaXMuX2lkcy5nZXQodik7XG5cdFx0XHRcdGlmIChpZCA9PT0gdW5kZWZpbmVkKSB7XG5cdFx0XHRcdFx0aWQgPSB0aGlzLl9pZFBvb2wucG9wKCk7XG5cdFx0XHRcdFx0aWYgKGlkID09PSB1bmRlZmluZWQpIHtcblx0XHRcdFx0XHRcdGlkID0gdGhpcy5fdmFsdWVzLmxlbmd0aDtcblx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0dGhpcy5fdmFsdWVzW2lkXSA9IHY7XG5cdFx0XHRcdFx0dGhpcy5fZ29SZWZDb3VudHNbaWRdID0gMDtcblx0XHRcdFx0XHR0aGlzLl9pZHMuc2V0KHYsIGlkKTtcblx0XHRcdFx0fVxuXHRcdFx0XHR0aGlzLl9nb1JlZkNvdW50c1tpZF0rKztcblx0XHRcdFx0bGV0IHR5cGVGbGFnID0gMDtcblx0XHRcdFx0c3dpdGNoICh0eXBlb2Ygdikge1xuXHRcdFx0XHRcdGNhc2UgXCJvYmplY3RcIjpcblx0XHRcdFx0XHRcdGlmICh2ICE9PSBudWxsKSB7XG5cdFx0XHRcdFx0XHRcdHR5cGVGbGFnID0gMTtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHRcdGJyZWFrO1xuXHRcdFx0XHRcdGNhc2UgXCJzdHJpbmdcIjpcblx0XHRcdFx0XHRcdHR5cGVGbGFnID0gMjtcblx0XHRcdFx0XHRcdGJyZWFrO1xuXHRcdFx0XHRcdGNhc2UgXCJzeW1ib2xcIjpcblx0XHRcdFx0XHRcdHR5cGVGbGFnID0gMztcblx0XHRcdFx0XHRcdGJyZWFrO1xuXHRcdFx0XHRcdGNhc2UgXCJmdW5jdGlvblwiOlxuXHRcdFx0XHRcdFx0dHlwZUZsYWcgPSA0O1xuXHRcdFx0XHRcdFx0YnJlYWs7XG5cdFx0XHRcdH1cblx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDMyKGFkZHIgKyA0LCBuYW5IZWFkIHwgdHlwZUZsYWcsIHRydWUpO1xuXHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIoYWRkciwgaWQsIHRydWUpO1xuXHRcdFx0fVxuXG5cdFx0XHRjb25zdCBsb2FkU2xpY2UgPSAoYWRkcikgPT4ge1xuXHRcdFx0XHRjb25zdCBhcnJheSA9IGdldEludDY0KGFkZHIgKyAwKTtcblx0XHRcdFx0Y29uc3QgbGVuID0gZ2V0SW50NjQoYWRkciArIDgpO1xuXHRcdFx0XHRyZXR1cm4gbmV3IFVpbnQ4QXJyYXkodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIGFycmF5LCBsZW4pO1xuXHRcdFx0fVxuXG5cdFx0XHRjb25zdCBsb2FkU2xpY2VPZlZhbHVlcyA9IChhZGRyKSA9PiB7XG5cdFx0XHRcdGNvbnN0IGFycmF5ID0gZ2V0SW50NjQoYWRkciArIDApO1xuXHRcdFx0XHRjb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7XG5cdFx0XHRcdGNvbnN0IGEgPSBuZXcgQXJyYXkobGVuKTtcblx0XHRcdFx0Zm9yIChsZXQgaSA9IDA7IGkgPCBsZW47IGkrKykge1xuXHRcdFx0XHRcdGFbaV0gPSBsb2FkVmFsdWUoYXJyYXkgKyBpICogOCk7XG5cdFx0XHRcdH1cblx0XHRcdFx0cmV0dXJuIGE7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IGxvYWRTdHJpbmcgPSAoYWRkcikgPT4ge1xuXHRcdFx0XHRjb25zdCBzYWRkciA9IGdldEludDY0KGFkZHIgKyAwKTtcblx0XHRcdFx0Y29uc3QgbGVuID0gZ2V0SW50NjQoYWRkciArIDgpO1xuXHRcdFx0XHRyZXR1cm4gZGVjb2Rlci5kZWNvZGUobmV3IERhdGFWaWV3KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLCBzYWRkciwgbGVuKSk7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IHRpbWVPcmlnaW4gPSBEYXRlLm5vdygpIC0gcGVyZm9ybWFuY2Uubm93KCk7XG5cdFx0XHR0aGlzLmltcG9ydE9iamVjdCA9IHtcblx0XHRcdFx0X2dvdGVzdDoge1xuXHRcdFx0XHRcdGFkZDogKGEsIGIpID0+IGEgKyBiLFxuXHRcdFx0XHR9LFxuXHRcdFx0XHRnb2pzOiB7XG5cdFx0XHRcdFx0Ly8gR28ncyBTUCBkb2VzIG5vdCBjaGFuZ2UgYXMgbG9uZyBhcyBubyBHbyBjb2RlIGlzIHJ1bm5pbmcuIFNvbWUgb3BlcmF0aW9ucyAoZS5nLiBjYWxscywgZ2V0dGVycyBhbmQgc2V0dGVycylcblx0XHRcdFx0XHQvLyBtYXkgc3luY2hyb25vdXNseSB0cmlnZ2VyIGEgR28gZXZlbnQgaGFuZGxlci4gVGhpcyBtYWtlcyBHbyBjb2RlIGdldCBleGVjdXRlZCBpbiB0aGUgbWlkZGxlIG9mIHRoZSBpbXBvcnRlZFxuXHRcdFx0XHRcdC8vIGZ1bmN0aW9uLiBBIGdvcm91dGluZSBjYW4gc3dpdGNoIHRvIGEgbmV3IHN0YWNrIGlmIHRoZSBjdXJyZW50IHN0YWNrIGlzIHRvbyBzbWFsbCAoc2VlIG1vcmVzdGFjayBmdW5jdGlvbikuXG5cdFx0XHRcdFx0Ly8gVGhpcyBjaGFuZ2VzIHRoZSBTUCwgdGh1cyB3ZSBoYXZlIHRvIHVwZGF0ZSB0aGUgU1AgdXNlZCBieSB0aGUgaW1wb3J0ZWQgZnVuY3Rpb24uXG5cblx0XHRcdFx0XHQvLyBmdW5jIHdhc21FeGl0KGNvZGUgaW50MzIpXG5cdFx0XHRcdFx0XCJydW50aW1lLndhc21FeGl0XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgY29kZSA9IHRoaXMubWVtLmdldEludDMyKHNwICsgOCwgdHJ1ZSk7XG5cdFx0XHRcdFx0XHR0aGlzLmV4aXRlZCA9IHRydWU7XG5cdFx0XHRcdFx0XHRkZWxldGUgdGhpcy5faW5zdDtcblx0XHRcdFx0XHRcdGRlbGV0ZSB0aGlzLl92YWx1ZXM7XG5cdFx0XHRcdFx0XHRkZWxldGUgdGhpcy5fZ29SZWZDb3VudHM7XG5cdFx0XHRcdFx0XHRkZWxldGUgdGhpcy5faWRzO1xuXHRcdFx0XHRcdFx0ZGVsZXRlIHRoaXMuX2lkUG9vbDtcblx0XHRcdFx0XHRcdHRoaXMuZXhpdChjb2RlKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB3YXNtV3JpdGUoZmQgdWludHB0ciwgcCB1bnNhZmUuUG9pbnRlciwgbiBpbnQzMilcblx0XHRcdFx0XHRcInJ1bnRpbWUud2FzbVdyaXRlXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgZmQgPSBnZXRJbnQ2NChzcCArIDgpO1xuXHRcdFx0XHRcdFx0Y29uc3QgcCA9IGdldEludDY0KHNwICsgMTYpO1xuXHRcdFx0XHRcdFx0Y29uc3QgbiA9IHRoaXMubWVtLmdldEludDMyKHNwICsgMjQsIHRydWUpO1xuXHRcdFx0XHRcdFx0ZnMud3JpdGVTeW5jKGZkLCBuZXcgVWludDhBcnJheSh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlciwgcCwgbikpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHJlc2V0TWVtb3J5RGF0YVZpZXcoKVxuXHRcdFx0XHRcdFwicnVudGltZS5yZXNldE1lbW9yeURhdGFWaWV3XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0dGhpcy5tZW0gPSBuZXcgRGF0YVZpZXcodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIG5hbm90aW1lMSgpIGludDY0XG5cdFx0XHRcdFx0XCJydW50aW1lLm5hbm90aW1lMVwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHNldEludDY0KHNwICsgOCwgKHRpbWVPcmlnaW4gKyBwZXJmb3JtYW5jZS5ub3coKSkgKiAxMDAwMDAwKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB3YWxsdGltZSgpIChzZWMgaW50NjQsIG5zZWMgaW50MzIpXG5cdFx0XHRcdFx0XCJydW50aW1lLndhbGx0aW1lXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgbXNlYyA9IChuZXcgRGF0ZSkuZ2V0VGltZSgpO1xuXHRcdFx0XHRcdFx0c2V0SW50NjQoc3AgKyA4LCBtc2VjIC8gMTAwMCk7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRJbnQzMihzcCArIDE2LCAobXNlYyAlIDEwMDApICogMTAwMDAwMCwgdHJ1ZSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgc2NoZWR1bGVUaW1lb3V0RXZlbnQoZGVsYXkgaW50NjQpIGludDMyXG5cdFx0XHRcdFx0XCJydW50aW1lLnNjaGVkdWxlVGltZW91dEV2ZW50XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgaWQgPSB0aGlzLl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQ7XG5cdFx0XHRcdFx0XHR0aGlzLl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQrKztcblx0XHRcdFx0XHRcdHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLnNldChpZCwgc2V0VGltZW91dChcblx0XHRcdFx0XHRcdFx0KCkgPT4ge1xuXHRcdFx0XHRcdFx0XHRcdHRoaXMuX3Jlc3VtZSgpO1xuXHRcdFx0XHRcdFx0XHRcdHdoaWxlICh0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cy5oYXMoaWQpKSB7XG5cdFx0XHRcdFx0XHRcdFx0XHQvLyBmb3Igc29tZSByZWFzb24gR28gZmFpbGVkIHRvIHJlZ2lzdGVyIHRoZSB0aW1lb3V0IGV2ZW50LCBsb2cgYW5kIHRyeSBhZ2FpblxuXHRcdFx0XHRcdFx0XHRcdFx0Ly8gKHRlbXBvcmFyeSB3b3JrYXJvdW5kIGZvciBodHRwczovL2dpdGh1Yi5jb20vZ29sYW5nL2dvL2lzc3Vlcy8yODk3NSlcblx0XHRcdFx0XHRcdFx0XHRcdGNvbnNvbGUud2FybihcInNjaGVkdWxlVGltZW91dEV2ZW50OiBtaXNzZWQgdGltZW91dCBldmVudFwiKTtcblx0XHRcdFx0XHRcdFx0XHRcdHRoaXMuX3Jlc3VtZSgpO1xuXHRcdFx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHRcdFx0fSxcblx0XHRcdFx0XHRcdFx0Z2V0SW50NjQoc3AgKyA4KSxcblx0XHRcdFx0XHRcdCkpO1xuXHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0SW50MzIoc3AgKyAxNiwgaWQsIHRydWUpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIGNsZWFyVGltZW91dEV2ZW50KGlkIGludDMyKVxuXHRcdFx0XHRcdFwicnVudGltZS5jbGVhclRpbWVvdXRFdmVudFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdGNvbnN0IGlkID0gdGhpcy5tZW0uZ2V0SW50MzIoc3AgKyA4LCB0cnVlKTtcblx0XHRcdFx0XHRcdGNsZWFyVGltZW91dCh0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cy5nZXQoaWQpKTtcblx0XHRcdFx0XHRcdHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLmRlbGV0ZShpZCk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgZ2V0UmFuZG9tRGF0YShyIFtdYnl0ZSlcblx0XHRcdFx0XHRcInJ1bnRpbWUuZ2V0UmFuZG9tRGF0YVwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdGNyeXB0by5nZXRSYW5kb21WYWx1ZXMobG9hZFNsaWNlKHNwICsgOCkpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIGZpbmFsaXplUmVmKHYgcmVmKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy5maW5hbGl6ZVJlZlwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdGNvbnN0IGlkID0gdGhpcy5tZW0uZ2V0VWludDMyKHNwICsgOCwgdHJ1ZSk7XG5cdFx0XHRcdFx0XHR0aGlzLl9nb1JlZkNvdW50c1tpZF0tLTtcblx0XHRcdFx0XHRcdGlmICh0aGlzLl9nb1JlZkNvdW50c1tpZF0gPT09IDApIHtcblx0XHRcdFx0XHRcdFx0Y29uc3QgdiA9IHRoaXMuX3ZhbHVlc1tpZF07XG5cdFx0XHRcdFx0XHRcdHRoaXMuX3ZhbHVlc1tpZF0gPSBudWxsO1xuXHRcdFx0XHRcdFx0XHR0aGlzLl9pZHMuZGVsZXRlKHYpO1xuXHRcdFx0XHRcdFx0XHR0aGlzLl9pZFBvb2wucHVzaChpZCk7XG5cdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgc3RyaW5nVmFsKHZhbHVlIHN0cmluZykgcmVmXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnN0cmluZ1ZhbFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyAyNCwgbG9hZFN0cmluZyhzcCArIDgpKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZUdldCh2IHJlZiwgcCBzdHJpbmcpIHJlZlxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUdldFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdGNvbnN0IHJlc3VsdCA9IFJlZmxlY3QuZ2V0KGxvYWRWYWx1ZShzcCArIDgpLCBsb2FkU3RyaW5nKHNwICsgMTYpKTtcblx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgMzIsIHJlc3VsdCk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVTZXQodiByZWYsIHAgc3RyaW5nLCB4IHJlZilcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVTZXRcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRSZWZsZWN0LnNldChsb2FkVmFsdWUoc3AgKyA4KSwgbG9hZFN0cmluZyhzcCArIDE2KSwgbG9hZFZhbHVlKHNwICsgMzIpKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZURlbGV0ZSh2IHJlZiwgcCBzdHJpbmcpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlRGVsZXRlXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0UmVmbGVjdC5kZWxldGVQcm9wZXJ0eShsb2FkVmFsdWUoc3AgKyA4KSwgbG9hZFN0cmluZyhzcCArIDE2KSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVJbmRleCh2IHJlZiwgaSBpbnQpIHJlZlxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUluZGV4XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDI0LCBSZWZsZWN0LmdldChsb2FkVmFsdWUoc3AgKyA4KSwgZ2V0SW50NjQoc3AgKyAxNikpKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gdmFsdWVTZXRJbmRleCh2IHJlZiwgaSBpbnQsIHggcmVmKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZVNldEluZGV4XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0UmVmbGVjdC5zZXQobG9hZFZhbHVlKHNwICsgOCksIGdldEludDY0KHNwICsgMTYpLCBsb2FkVmFsdWUoc3AgKyAyNCkpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHZhbHVlQ2FsbCh2IHJlZiwgbSBzdHJpbmcsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlQ2FsbFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHRyeSB7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IHYgPSBsb2FkVmFsdWUoc3AgKyA4KTtcblx0XHRcdFx0XHRcdFx0Y29uc3QgbSA9IFJlZmxlY3QuZ2V0KHYsIGxvYWRTdHJpbmcoc3AgKyAxNikpO1xuXHRcdFx0XHRcdFx0XHRjb25zdCBhcmdzID0gbG9hZFNsaWNlT2ZWYWx1ZXMoc3AgKyAzMik7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IHJlc3VsdCA9IFJlZmxlY3QuYXBwbHkobSwgdiwgYXJncyk7XG5cdFx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyA1NiwgcmVzdWx0KTtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA2NCwgMSk7XG5cdFx0XHRcdFx0XHR9IGNhdGNoIChlcnIpIHtcblx0XHRcdFx0XHRcdFx0c3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmVcblx0XHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDU2LCBlcnIpO1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDY0LCAwKTtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZUludm9rZSh2IHJlZiwgYXJncyBbXXJlZikgKHJlZiwgYm9vbClcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVJbnZva2VcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHR0cnkge1xuXHRcdFx0XHRcdFx0XHRjb25zdCB2ID0gbG9hZFZhbHVlKHNwICsgOCk7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDE2KTtcblx0XHRcdFx0XHRcdFx0Y29uc3QgcmVzdWx0ID0gUmVmbGVjdC5hcHBseSh2LCB1bmRlZmluZWQsIGFyZ3MpO1xuXHRcdFx0XHRcdFx0XHRzcCA9IHRoaXMuX2luc3QuZXhwb3J0cy5nZXRzcCgpID4+PiAwOyAvLyBzZWUgY29tbWVudCBhYm92ZVxuXHRcdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgNDAsIHJlc3VsdCk7XG5cdFx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDEpO1xuXHRcdFx0XHRcdFx0fSBjYXRjaCAoZXJyKSB7XG5cdFx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyA0MCwgZXJyKTtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMCk7XG5cdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVOZXcodiByZWYsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlTmV3XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0dHJ5IHtcblx0XHRcdFx0XHRcdFx0Y29uc3QgdiA9IGxvYWRWYWx1ZShzcCArIDgpO1xuXHRcdFx0XHRcdFx0XHRjb25zdCBhcmdzID0gbG9hZFNsaWNlT2ZWYWx1ZXMoc3AgKyAxNik7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IHJlc3VsdCA9IFJlZmxlY3QuY29uc3RydWN0KHYsIGFyZ3MpO1xuXHRcdFx0XHRcdFx0XHRzcCA9IHRoaXMuX2luc3QuZXhwb3J0cy5nZXRzcCgpID4+PiAwOyAvLyBzZWUgY29tbWVudCBhYm92ZVxuXHRcdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgNDAsIHJlc3VsdCk7XG5cdFx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDEpO1xuXHRcdFx0XHRcdFx0fSBjYXRjaCAoZXJyKSB7XG5cdFx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyA0MCwgZXJyKTtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMCk7XG5cdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVMZW5ndGgodiByZWYpIGludFxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUxlbmd0aFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHNldEludDY0KHNwICsgMTYsIHBhcnNlSW50KGxvYWRWYWx1ZShzcCArIDgpLmxlbmd0aCkpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyB2YWx1ZVByZXBhcmVTdHJpbmcodiByZWYpIChyZWYsIGludClcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVQcmVwYXJlU3RyaW5nXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3Qgc3RyID0gZW5jb2Rlci5lbmNvZGUoU3RyaW5nKGxvYWRWYWx1ZShzcCArIDgpKSk7XG5cdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgMTYsIHN0cik7XG5cdFx0XHRcdFx0XHRzZXRJbnQ2NChzcCArIDI0LCBzdHIubGVuZ3RoKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gdmFsdWVMb2FkU3RyaW5nKHYgcmVmLCBiIFtdYnl0ZSlcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVMb2FkU3RyaW5nXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3Qgc3RyID0gbG9hZFZhbHVlKHNwICsgOCk7XG5cdFx0XHRcdFx0XHRsb2FkU2xpY2Uoc3AgKyAxNikuc2V0KHN0cik7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVJbnN0YW5jZU9mKHYgcmVmLCB0IHJlZikgYm9vbFxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUluc3RhbmNlT2ZcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDI0LCAobG9hZFZhbHVlKHNwICsgOCkgaW5zdGFuY2VvZiBsb2FkVmFsdWUoc3AgKyAxNikpID8gMSA6IDApO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIGNvcHlCeXRlc1RvR28oZHN0IFtdYnl0ZSwgc3JjIHJlZikgKGludCwgYm9vbClcblx0XHRcdFx0XHRcInN5c2NhbGwvanMuY29weUJ5dGVzVG9Hb1wiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdGNvbnN0IGRzdCA9IGxvYWRTbGljZShzcCArIDgpO1xuXHRcdFx0XHRcdFx0Y29uc3Qgc3JjID0gbG9hZFZhbHVlKHNwICsgMzIpO1xuXHRcdFx0XHRcdFx0aWYgKCEoc3JjIGluc3RhbmNlb2YgVWludDhBcnJheSB8fCBzcmMgaW5zdGFuY2VvZiBVaW50OENsYW1wZWRBcnJheSkpIHtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMCk7XG5cdFx0XHRcdFx0XHRcdHJldHVybjtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHRcdGNvbnN0IHRvQ29weSA9IHNyYy5zdWJhcnJheSgwLCBkc3QubGVuZ3RoKTtcblx0XHRcdFx0XHRcdGRzdC5zZXQodG9Db3B5KTtcblx0XHRcdFx0XHRcdHNldEludDY0KHNwICsgNDAsIHRvQ29weS5sZW5ndGgpO1xuXHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgY29weUJ5dGVzVG9KUyhkc3QgcmVmLCBzcmMgW11ieXRlKSAoaW50LCBib29sKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy5jb3B5Qnl0ZXNUb0pTXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgZHN0ID0gbG9hZFZhbHVlKHNwICsgOCk7XG5cdFx0XHRcdFx0XHRjb25zdCBzcmMgPSBsb2FkU2xpY2Uoc3AgKyAxNik7XG5cdFx0XHRcdFx0XHRpZiAoIShkc3QgaW5zdGFuY2VvZiBVaW50OEFycmF5IHx8IGRzdCBpbnN0YW5jZW9mIFVpbnQ4Q2xhbXBlZEFycmF5KSkge1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTtcblx0XHRcdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHRcdFx0fVxuXHRcdFx0XHRcdFx0Y29uc3QgdG9Db3B5ID0gc3JjLnN1YmFycmF5KDAsIGRzdC5sZW5ndGgpO1xuXHRcdFx0XHRcdFx0ZHN0LnNldCh0b0NvcHkpO1xuXHRcdFx0XHRcdFx0c2V0SW50NjQoc3AgKyA0MCwgdG9Db3B5Lmxlbmd0aCk7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAxKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0XCJkZWJ1Z1wiOiAodmFsdWUpID0+IHtcblx0XHRcdFx0XHRcdGNvbnNvbGUubG9nKHZhbHVlKTtcblx0XHRcdFx0XHR9LFxuXHRcdFx0XHR9XG5cdFx0XHR9O1xuXHRcdH1cblxuXHRcdGFzeW5jIHJ1bihpbnN0YW5jZSkge1xuXHRcdFx0aWYgKCEoaW5zdGFuY2UgaW5zdGFuY2VvZiBXZWJBc3NlbWJseS5JbnN0YW5jZSkpIHtcblx0XHRcdFx0dGhyb3cgbmV3IEVycm9yKFwiR28ucnVuOiBXZWJBc3NlbWJseS5JbnN0YW5jZSBleHBlY3RlZFwiKTtcblx0XHRcdH1cblx0XHRcdHRoaXMuX2luc3QgPSBpbnN0YW5jZTtcblx0XHRcdHRoaXMubWVtID0gbmV3IERhdGFWaWV3KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyKTtcblx0XHRcdHRoaXMuX3ZhbHVlcyA9IFsgLy8gSlMgdmFsdWVzIHRoYXQgR28gY3VycmVudGx5IGhhcyByZWZlcmVuY2VzIHRvLCBpbmRleGVkIGJ5IHJlZmVyZW5jZSBpZFxuXHRcdFx0XHROYU4sXG5cdFx0XHRcdDAsXG5cdFx0XHRcdG51bGwsXG5cdFx0XHRcdHRydWUsXG5cdFx0XHRcdGZhbHNlLFxuXHRcdFx0XHRnbG9iYWxUaGlzLFxuXHRcdFx0XHR0aGlzLFxuXHRcdFx0XTtcblx0XHRcdHRoaXMuX2dvUmVmQ291bnRzID0gbmV3IEFycmF5KHRoaXMuX3ZhbHVlcy5sZW5ndGgpLmZpbGwoSW5maW5pdHkpOyAvLyBudW1iZXIgb2YgcmVmZXJlbmNlcyB0aGF0IEdvIGhhcyB0byBhIEpTIHZhbHVlLCBpbmRleGVkIGJ5IHJlZmVyZW5jZSBpZFxuXHRcdFx0dGhpcy5faWRzID0gbmV3IE1hcChbIC8vIG1hcHBpbmcgZnJvbSBKUyB2YWx1ZXMgdG8gcmVmZXJlbmNlIGlkc1xuXHRcdFx0XHRbMCwgMV0sXG5cdFx0XHRcdFtudWxsLCAyXSxcblx0XHRcdFx0W3RydWUsIDNdLFxuXHRcdFx0XHRbZmFsc2UsIDRdLFxuXHRcdFx0XHRbZ2xvYmFsVGhpcywgNV0sXG5cdFx0XHRcdFt0aGlzLCA2XSxcblx0XHRcdF0pO1xuXHRcdFx0dGhpcy5faWRQb29sID0gW107ICAgLy8gdW51c2VkIGlkcyB0aGF0IGhhdmUgYmVlbiBnYXJiYWdlIGNvbGxlY3RlZFxuXHRcdFx0dGhpcy5leGl0ZWQgPSBmYWxzZTsgLy8gd2hldGhlciB0aGUgR28gcHJvZ3JhbSBoYXMgZXhpdGVkXG5cblx0XHRcdC8vIFBhc3MgY29tbWFuZCBsaW5lIGFyZ3VtZW50cyBhbmQgZW52aXJvbm1lbnQgdmFyaWFibGVzIHRvIFdlYkFzc2VtYmx5IGJ5IHdyaXRpbmcgdGhlbSB0byB0aGUgbGluZWFyIG1lbW9yeS5cblx0XHRcdGxldCBvZmZzZXQgPSA0MDk2O1xuXG5cdFx0XHRjb25zdCBzdHJQdHIgPSAoc3RyKSA9PiB7XG5cdFx0XHRcdGNvbnN0IHB0ciA9IG9mZnNldDtcblx0XHRcdFx0Y29uc3QgYnl0ZXMgPSBlbmNvZGVyLmVuY29kZShzdHIgKyBcIlxcMFwiKTtcblx0XHRcdFx0bmV3IFVpbnQ4QXJyYXkodGhpcy5tZW0uYnVmZmVyLCBvZmZzZXQsIGJ5dGVzLmxlbmd0aCkuc2V0KGJ5dGVzKTtcblx0XHRcdFx0b2Zmc2V0ICs9IGJ5dGVzLmxlbmd0aDtcblx0XHRcdFx0aWYgKG9mZnNldCAlIDggIT09IDApIHtcblx0XHRcdFx0XHRvZmZzZXQgKz0gOCAtIChvZmZzZXQgJSA4KTtcblx0XHRcdFx0fVxuXHRcdFx0XHRyZXR1cm4gcHRyO1xuXHRcdFx0fTtcblxuXHRcdFx0Y29uc3QgYXJnYyA9IHRoaXMuYXJndi5sZW5ndGg7XG5cblx0XHRcdGNvbnN0IGFyZ3ZQdHJzID0gW107XG5cdFx0XHR0aGlzLmFyZ3YuZm9yRWFjaCgoYXJnKSA9PiB7XG5cdFx0XHRcdGFyZ3ZQdHJzLnB1c2goc3RyUHRyKGFyZykpO1xuXHRcdFx0fSk7XG5cdFx0XHRhcmd2UHRycy5wdXNoKDApO1xuXG5cdFx0XHRjb25zdCBrZXlzID0gT2JqZWN0LmtleXModGhpcy5lbnYpLnNvcnQoKTtcblx0XHRcdGtleXMuZm9yRWFjaCgoa2V5KSA9PiB7XG5cdFx0XHRcdGFyZ3ZQdHJzLnB1c2goc3RyUHRyKGAke2tleX09JHt0aGlzLmVudltrZXldfWApKTtcblx0XHRcdH0pO1xuXHRcdFx0YXJndlB0cnMucHVzaCgwKTtcblxuXHRcdFx0Y29uc3QgYXJndiA9IG9mZnNldDtcblx0XHRcdGFyZ3ZQdHJzLmZvckVhY2goKHB0cikgPT4ge1xuXHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIob2Zmc2V0LCBwdHIsIHRydWUpO1xuXHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIob2Zmc2V0ICsgNCwgMCwgdHJ1ZSk7XG5cdFx0XHRcdG9mZnNldCArPSA4O1xuXHRcdFx0fSk7XG5cblx0XHRcdC8vIFRoZSBsaW5rZXIgZ3VhcmFudGVlcyBnbG9iYWwgZGF0YSBzdGFydHMgZnJvbSBhdCBsZWFzdCB3YXNtTWluRGF0YUFkZHIuXG5cdFx0XHQvLyBLZWVwIGluIHN5bmMgd2l0aCBjbWQvbGluay9pbnRlcm5hbC9sZC9kYXRhLmdvOndhc21NaW5EYXRhQWRkci5cblx0XHRcdGNvbnN0IHdhc21NaW5EYXRhQWRkciA9IDQwOTYgKyA4MTkyO1xuXHRcdFx0aWYgKG9mZnNldCA+PSB3YXNtTWluRGF0YUFkZHIpIHtcblx0XHRcdFx0dGhyb3cgbmV3IEVycm9yKFwidG90YWwgbGVuZ3RoIG9mIGNvbW1hbmQgbGluZSBhbmQgZW52aXJvbm1lbnQgdmFyaWFibGVzIGV4Y2VlZHMgbGltaXRcIik7XG5cdFx0XHR9XG5cblx0XHRcdHRoaXMuX2luc3QuZXhwb3J0cy5ydW4oYXJnYywgYXJndik7XG5cdFx0XHRpZiAodGhpcy5leGl0ZWQpIHtcblx0XHRcdFx0dGhpcy5fcmVzb2x2ZUV4aXRQcm9taXNlKCk7XG5cdFx0XHR9XG5cdFx0XHRhd2FpdCB0aGlzLl9leGl0UHJvbWlzZTtcblx0XHR9XG5cblx0XHRfcmVzdW1lKCkge1xuXHRcdFx0aWYgKHRoaXMuZXhpdGVkKSB7XG5cdFx0XHRcdHRocm93IG5ldyBFcnJvcihcIkdvIHByb2dyYW0gaGFzIGFscmVhZHkgZXhpdGVkXCIpO1xuXHRcdFx0fVxuXHRcdFx0dGhpcy5faW5zdC5leHBvcnRzLnJlc3VtZSgpO1xuXHRcdFx0aWYgKHRoaXMuZXhpdGVkKSB7XG5cdFx0XHRcdHRoaXMuX3Jlc29sdmVFeGl0UHJvbWlzZSgpO1xuXHRcdFx0fVxuXHRcdH1cblxuXHRcdF9tYWtlRnVuY1dyYXBwZXIoaWQpIHtcblx0XHRcdGNvbnN0IGdvID0gdGhpcztcblx0XHRcdHJldHVybiBmdW5jdGlvbiAoKSB7XG5cdFx0XHRcdGNvbnN0IGV2ZW50ID0geyBpZDogaWQsIHRoaXM6IHRoaXMsIGFyZ3M6IGFyZ3VtZW50cyB9O1xuXHRcdFx0XHRnby5fcGVuZGluZ0V2ZW50ID0gZXZlbnQ7XG5cdFx0XHRcdGdvLl9yZXN1bWUoKTtcblx0XHRcdFx0cmV0dXJuIGV2ZW50LnJlc3VsdDtcblx0XHRcdH07XG5cdFx0fVxuXHR9XG59KSgpO1xuXG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL3ZpZXdzL3dhc21fZXhlYy5qcyJdLCJzb3VyY2VSb290IjoiIn0=