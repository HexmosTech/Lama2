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
  var l2codeElements = document.querySelectorAll(".".concat(className));
  l2codeElements.forEach(function (divElement) {
    var l2content = divElement.textContent;
    divElement.textContent = '';
    var elementType = divElement.tagName; // Get the tag name of the divElement
    console.log('Element type:', elementType);
    var textareaElement = document.createElement('textarea');
    textareaElement.id = 'myTextarea';
    textareaElement.name = 'Hash';
    textareaElement.style.maxWidth = '100%'; // Set a maximum width to make it responsive
    textareaElement.style.width = '500px';
    textareaElement.style.height = '200px';
    textareaElement.style.backgroundColor = '#f0f0f0';
    textareaElement.style.padding = '10px';
    textareaElement.textContent = l2content;
    divElement.appendChild(textareaElement);
    var buttonImage = document.createElement('img');
    buttonImage.src = 'arrowl2.png';
    buttonImage.style.width = '30px';
    var index = Math.random().toString(36).substr(2, 10);
    var buttonElement = document.createElement('button');
    buttonElement.appendChild(buttonImage);
    buttonElement.addEventListener('click', function () {
      return addTextToDiv(divElement, index, go);
    });
    buttonElement.style.all = 'unset';
    buttonElement.style.position = 'relative';
    buttonElement.style.bottom = '150px';
    buttonElement.style.right = '50px';
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
  });
}
function addTextToDiv(_x, _x2, _x3) {
  return _addTextToDiv.apply(this, arguments);
}
function _addTextToDiv() {
  _addTextToDiv = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee(divElement, index, go) {
    var textArea, outputId, existingOutputDiv, outputDiv;
    return _regeneratorRuntime().wrap(function _callee$(_context) {
      while (1) switch (_context.prev = _context.next) {
        case 0:
          textArea = divElement.querySelector('textarea');
          outputId = "outputField_".concat(index);
          existingOutputDiv = document.getElementById(outputId);
          if (textArea) {
            console.log("Text inside the text area:", textArea.value);
            if (!existingOutputDiv) {
              console.log("Text inside the text area:", textArea.value);
              outputDiv = document.createElement('div');
              outputDiv.className = 'outputField';
              outputDiv.id = "outputField_".concat(index);
              outputDiv.style.fontSize = '16px';
              outputDiv.style.marginTop = '10px';
              outputDiv.style.width = '500px';
              outputDiv.style.minHeight = '200px';
              outputDiv.style.color = 'white';
              outputDiv.style.height = '400px';
              outputDiv.style.overflowY = 'auto';
              outputDiv.style.backgroundColor = '#333';
              outputDiv.style.border = '1px solid #555';
              outputDiv.style.borderRadius = '5px';
              outputDiv.style.padding = '20px';
              outputDiv.style.fontFamily = 'Arial, sans-serif';
              outputDiv.style.overflowY = 'auto';
              outputDiv.style.whiteSpace = 'pre-wrap';
              outputDiv.style.boxShadow = '0 0 10px rgba(0, 0, 0, 0.2)';
              outputDiv.style.transition = 'background-color 0.3s';
              divElement.appendChild(outputDiv);
              // makeLamaRequest(textArea.value, divElement, index);
            }

            makeLamaRequest(textArea.value, divElement, index, go);
          } else {
            console.log("No text area found. Logging data from divElement:", divElement.textContent);
            makeLamaRequest(divElement.textContent, divElement, index);
          }
        case 4:
        case "end":
          return _context.stop();
      }
    }, _callee);
  }));
  return _addTextToDiv.apply(this, arguments);
}
function makeLamaRequest(_x4, _x5, _x6, _x7) {
  return _makeLamaRequest.apply(this, arguments);
}
function _makeLamaRequest() {
  _makeLamaRequest = _asyncToGenerator( /*#__PURE__*/_regeneratorRuntime().mark(function _callee2(command, divElement, index, go) {
    var response, jsonObject, formattedBody, respStatusCode, formattedJson, outputField, statusCodeDiv;
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
          console.log("go object :", go);
          _context2.next = 7;
          return goWebRequestFunc(command);
        case 7:
          response = _context2.sent;
          console.log("response Data:", response);
          jsonObject = JSON.parse(response);
          formattedBody = JSON.parse(jsonObject.Body);
          respStatusCode = JSON.parse(jsonObject.StatusCode);
          formattedJson = JSON.stringify(formattedBody, null, 2); // var formattedJson = JSON.stringify(jsonObject, null, 2);
          console.log("formattedJson:", formattedJson);
          outputField = divElement.querySelector("#outputField_".concat(index));
          console.log(outputField);
          outputField.textContent = formattedJson;
          console.log("outputField.textContent:", outputField.textContent);
          statusCodeDiv = document.createElement('div');
          statusCodeDiv.textContent = "Status Code: ".concat(respStatusCode);
          if (respStatusCode > 200) {
            statusCodeDiv.style.color = 'red'; // Set text color to red for status codes > 200
          } else {
            statusCodeDiv.style.color = 'green'; // Set text color to green for status codes <= 200
          }
          // statusCodeDiv.style.color = 'green'; // Set text color to green
          statusCodeDiv.style.fontWeight = 'bold'; // Make the text bold
          // divElement.appendChild(statusCodeDiv);
          divElement.insertBefore(statusCodeDiv, outputField);
          _context2.next = 28;
          break;
        case 25:
          _context2.prev = 25;
          _context2.t0 = _context2["catch"](0);
          console.error("Error making Lama request:", _context2.t0);
        case 28:
        case "end":
          return _context2.stop();
      }
    }, _callee2, null, [[0, 25]]);
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
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly8vd2VicGFjay9ib290c3RyYXAgYzA5NGZiOTkxMjgyMDA1MjcxNjAiLCJ3ZWJwYWNrOi8vLy4vc3JjL21haW4uanMiLCJ3ZWJwYWNrOi8vLy4vc3JjL3NlcnZpY2VzLmpzIiwid2VicGFjazovLy8uL3NyYy92aWV3cy9sYi5qcyIsIndlYnBhY2s6Ly8vLi9zcmMvdmlld3MvbGIuY3NzPzYzYTciLCJ3ZWJwYWNrOi8vLy4vc3JjL3ZpZXdzL2xiLmNzcyIsIndlYnBhY2s6Ly8vLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L3J1bnRpbWUvYXBpLmpzIiwid2VicGFjazovLy8uL25vZGVfbW9kdWxlcy9zdHlsZS1sb2FkZXIvbGliL2FkZFN0eWxlcy5qcyIsIndlYnBhY2s6Ly8vLi9ub2RlX21vZHVsZXMvc3R5bGUtbG9hZGVyL2xpYi91cmxzLmpzIiwid2VicGFjazovLy8uL3NyYy92aWV3cy93YXNtX2V4ZWMuanMiXSwibmFtZXMiOlsiX3JlZ2VuZXJhdG9yUnVudGltZSIsImUiLCJ0IiwiciIsIk9iamVjdCIsInByb3RvdHlwZSIsIm4iLCJoYXNPd25Qcm9wZXJ0eSIsIm8iLCJkZWZpbmVQcm9wZXJ0eSIsInZhbHVlIiwiaSIsIlN5bWJvbCIsImEiLCJpdGVyYXRvciIsImMiLCJhc3luY0l0ZXJhdG9yIiwidSIsInRvU3RyaW5nVGFnIiwiZGVmaW5lIiwiZW51bWVyYWJsZSIsImNvbmZpZ3VyYWJsZSIsIndyaXRhYmxlIiwid3JhcCIsIkdlbmVyYXRvciIsImNyZWF0ZSIsIkNvbnRleHQiLCJtYWtlSW52b2tlTWV0aG9kIiwidHJ5Q2F0Y2giLCJ0eXBlIiwiYXJnIiwiY2FsbCIsImgiLCJsIiwiZiIsInMiLCJ5IiwiR2VuZXJhdG9yRnVuY3Rpb24iLCJHZW5lcmF0b3JGdW5jdGlvblByb3RvdHlwZSIsInAiLCJkIiwiZ2V0UHJvdG90eXBlT2YiLCJ2IiwidmFsdWVzIiwiZyIsImRlZmluZUl0ZXJhdG9yTWV0aG9kcyIsImZvckVhY2giLCJfaW52b2tlIiwiQXN5bmNJdGVyYXRvciIsImludm9rZSIsIl90eXBlb2YiLCJyZXNvbHZlIiwiX19hd2FpdCIsInRoZW4iLCJjYWxsSW52b2tlV2l0aE1ldGhvZEFuZEFyZyIsIkVycm9yIiwiZG9uZSIsIm1ldGhvZCIsImRlbGVnYXRlIiwibWF5YmVJbnZva2VEZWxlZ2F0ZSIsInNlbnQiLCJfc2VudCIsImRpc3BhdGNoRXhjZXB0aW9uIiwiYWJydXB0IiwiVHlwZUVycm9yIiwicmVzdWx0TmFtZSIsIm5leHQiLCJuZXh0TG9jIiwicHVzaFRyeUVudHJ5IiwidHJ5TG9jIiwiY2F0Y2hMb2MiLCJmaW5hbGx5TG9jIiwiYWZ0ZXJMb2MiLCJ0cnlFbnRyaWVzIiwicHVzaCIsInJlc2V0VHJ5RW50cnkiLCJjb21wbGV0aW9uIiwicmVzZXQiLCJpc05hTiIsImxlbmd0aCIsImRpc3BsYXlOYW1lIiwiaXNHZW5lcmF0b3JGdW5jdGlvbiIsImNvbnN0cnVjdG9yIiwibmFtZSIsIm1hcmsiLCJzZXRQcm90b3R5cGVPZiIsIl9fcHJvdG9fXyIsImF3cmFwIiwiYXN5bmMiLCJQcm9taXNlIiwia2V5cyIsInJldmVyc2UiLCJwb3AiLCJwcmV2IiwiY2hhckF0Iiwic2xpY2UiLCJzdG9wIiwicnZhbCIsImhhbmRsZSIsImNvbXBsZXRlIiwiZmluaXNoIiwiX2NhdGNoIiwiZGVsZWdhdGVZaWVsZCIsImFzeW5jR2VuZXJhdG9yU3RlcCIsImdlbiIsInJlamVjdCIsIl9uZXh0IiwiX3Rocm93Iiwia2V5IiwiaW5mbyIsImVycm9yIiwiX2FzeW5jVG9HZW5lcmF0b3IiLCJmbiIsInNlbGYiLCJhcmdzIiwiYXJndW1lbnRzIiwiYXBwbHkiLCJlcnIiLCJ1bmRlZmluZWQiLCJzdXBwb3J0ZWRBUEkiLCJnbyIsImFwcCIsIndpbmRvdyIsImNvbnNvbGUiLCJsb2ciLCJjb25maWd1cmF0aW9ucyIsInNvbWVEZWZhdWx0Q29uZmlndXJhdGlvbiIsImdsb2JhbE9iamVjdCIsInF1ZXVlIiwicSIsInRvTG93ZXJDYXNlIiwiZXh0ZW5kT2JqZWN0IiwiYXBpSGFuZGxlciIsIl94IiwiX3gyIiwiX2FwaUhhbmRsZXIiLCJfY2FsbGVlIiwiYXBpIiwicGFyYW1zIiwiX2NhbGxlZSQiLCJfY29udGV4dCIsImluZGV4T2YiLCJjb25jYXQiLCJ0MCIsImxvYWRXYXNtIiwiYWRkQnV0dG9uVG9UZXh0YXJlYSIsIndhcm4iLCJiIiwiX2xvYWRXYXNtIiwiX2NhbGxlZTIiLCJyZXNwb25zZSIsImJ1ZmZlciIsIl9nbyIsInJlc3VsdCIsIl9jYWxsZWUyJCIsIl9jb250ZXh0MiIsImZldGNoIiwiYXJyYXlCdWZmZXIiLCJHbyIsIldlYkFzc2VtYmx5IiwiaW5zdGFudGlhdGUiLCJpbXBvcnRPYmplY3QiLCJydW4iLCJpbnN0YW5jZSIsInBpbmciLCJjbGFzc05hbWUiLCJsMmNvZGVFbGVtZW50cyIsImRvY3VtZW50IiwicXVlcnlTZWxlY3RvckFsbCIsImRpdkVsZW1lbnQiLCJsMmNvbnRlbnQiLCJ0ZXh0Q29udGVudCIsImVsZW1lbnRUeXBlIiwidGFnTmFtZSIsInRleHRhcmVhRWxlbWVudCIsImNyZWF0ZUVsZW1lbnQiLCJpZCIsInN0eWxlIiwibWF4V2lkdGgiLCJ3aWR0aCIsImhlaWdodCIsImJhY2tncm91bmRDb2xvciIsInBhZGRpbmciLCJhcHBlbmRDaGlsZCIsImJ1dHRvbkltYWdlIiwic3JjIiwiaW5kZXgiLCJNYXRoIiwicmFuZG9tIiwidG9TdHJpbmciLCJzdWJzdHIiLCJidXR0b25FbGVtZW50IiwiYWRkRXZlbnRMaXN0ZW5lciIsImFkZFRleHRUb0RpdiIsImFsbCIsInBvc2l0aW9uIiwiYm90dG9tIiwicmlnaHQiLCJ0cmFuc2l0aW9uIiwiY3Vyc29yIiwiaGlnaGxpZ2h0Q29sb3IiLCJsb2dvSW1hZ2UiLCJib3JkZXJSYWRpdXMiLCJsb2NhdGlvbiIsImhyZWYiLCJfeDMiLCJfYWRkVGV4dFRvRGl2IiwidGV4dEFyZWEiLCJvdXRwdXRJZCIsImV4aXN0aW5nT3V0cHV0RGl2Iiwib3V0cHV0RGl2IiwicXVlcnlTZWxlY3RvciIsImdldEVsZW1lbnRCeUlkIiwiZm9udFNpemUiLCJtYXJnaW5Ub3AiLCJtaW5IZWlnaHQiLCJjb2xvciIsIm92ZXJmbG93WSIsImJvcmRlciIsImZvbnRGYW1pbHkiLCJ3aGl0ZVNwYWNlIiwiYm94U2hhZG93IiwibWFrZUxhbWFSZXF1ZXN0IiwiX3g0IiwiX3g1IiwiX3g2IiwiX3g3IiwiX21ha2VMYW1hUmVxdWVzdCIsImNvbW1hbmQiLCJqc29uT2JqZWN0IiwiZm9ybWF0dGVkQm9keSIsInJlc3BTdGF0dXNDb2RlIiwiZm9ybWF0dGVkSnNvbiIsIm91dHB1dEZpZWxkIiwic3RhdHVzQ29kZURpdiIsImdvV2ViUmVxdWVzdEZ1bmMiLCJKU09OIiwicGFyc2UiLCJCb2R5IiwiU3RhdHVzQ29kZSIsInN0cmluZ2lmeSIsImZvbnRXZWlnaHQiLCJpbnNlcnRCZWZvcmUiLCJfY2xhc3NDYWxsQ2hlY2siLCJDb25zdHJ1Y3RvciIsIl9kZWZpbmVQcm9wZXJ0aWVzIiwidGFyZ2V0IiwicHJvcHMiLCJkZXNjcmlwdG9yIiwiX3RvUHJvcGVydHlLZXkiLCJfY3JlYXRlQ2xhc3MiLCJwcm90b1Byb3BzIiwic3RhdGljUHJvcHMiLCJfdG9QcmltaXRpdmUiLCJTdHJpbmciLCJpbnB1dCIsImhpbnQiLCJwcmltIiwidG9QcmltaXRpdmUiLCJyZXMiLCJOdW1iZXIiLCJlbm9zeXMiLCJjb2RlIiwiZ2xvYmFsVGhpcyIsImZzIiwib3V0cHV0QnVmIiwiY29uc3RhbnRzIiwiT19XUk9OTFkiLCJPX1JEV1IiLCJPX0NSRUFUIiwiT19UUlVOQyIsIk9fQVBQRU5EIiwiT19FWENMIiwid3JpdGVTeW5jIiwiZmQiLCJidWYiLCJkZWNvZGVyIiwiZGVjb2RlIiwibmwiLCJsYXN0SW5kZXhPZiIsInN1YnN0cmluZyIsIndyaXRlIiwib2Zmc2V0IiwiY2FsbGJhY2siLCJjaG1vZCIsInBhdGgiLCJtb2RlIiwiY2hvd24iLCJ1aWQiLCJnaWQiLCJjbG9zZSIsImZjaG1vZCIsImZjaG93biIsImZzdGF0IiwiZnN5bmMiLCJmdHJ1bmNhdGUiLCJsY2hvd24iLCJsaW5rIiwibHN0YXQiLCJta2RpciIsInBlcm0iLCJvcGVuIiwiZmxhZ3MiLCJyZWFkIiwicmVhZGRpciIsInJlYWRsaW5rIiwicmVuYW1lIiwiZnJvbSIsInRvIiwicm1kaXIiLCJzdGF0Iiwic3ltbGluayIsInRydW5jYXRlIiwidW5saW5rIiwidXRpbWVzIiwiYXRpbWUiLCJtdGltZSIsInByb2Nlc3MiLCJnZXR1aWQiLCJnZXRnaWQiLCJnZXRldWlkIiwiZ2V0ZWdpZCIsImdldGdyb3VwcyIsInBpZCIsInBwaWQiLCJ1bWFzayIsImN3ZCIsImNoZGlyIiwiY3J5cHRvIiwicGVyZm9ybWFuY2UiLCJUZXh0RW5jb2RlciIsIlRleHREZWNvZGVyIiwiZW5jb2RlciIsIl9jbGFzcyIsIl90aGlzIiwiYXJndiIsImVudiIsImV4aXQiLCJfZXhpdFByb21pc2UiLCJfcmVzb2x2ZUV4aXRQcm9taXNlIiwiX3BlbmRpbmdFdmVudCIsIl9zY2hlZHVsZWRUaW1lb3V0cyIsIk1hcCIsIl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQiLCJzZXRJbnQ2NCIsImFkZHIiLCJtZW0iLCJzZXRVaW50MzIiLCJmbG9vciIsInNldEludDMyIiwiZ2V0SW50NjQiLCJsb3ciLCJnZXRVaW50MzIiLCJoaWdoIiwiZ2V0SW50MzIiLCJsb2FkVmFsdWUiLCJnZXRGbG9hdDY0IiwiX3ZhbHVlcyIsInN0b3JlVmFsdWUiLCJuYW5IZWFkIiwic2V0RmxvYXQ2NCIsIl9pZHMiLCJnZXQiLCJfaWRQb29sIiwiX2dvUmVmQ291bnRzIiwic2V0IiwidHlwZUZsYWciLCJsb2FkU2xpY2UiLCJhcnJheSIsImxlbiIsIlVpbnQ4QXJyYXkiLCJfaW5zdCIsImV4cG9ydHMiLCJsb2FkU2xpY2VPZlZhbHVlcyIsIkFycmF5IiwibG9hZFN0cmluZyIsInNhZGRyIiwiRGF0YVZpZXciLCJ0aW1lT3JpZ2luIiwiRGF0ZSIsIm5vdyIsIl9nb3Rlc3QiLCJhZGQiLCJnb2pzIiwicnVudGltZVdhc21FeGl0Iiwic3AiLCJleGl0ZWQiLCJydW50aW1lV2FzbVdyaXRlIiwicnVudGltZVJlc2V0TWVtb3J5RGF0YVZpZXciLCJydW50aW1lTmFub3RpbWUxIiwicnVudGltZVdhbGx0aW1lIiwibXNlYyIsImdldFRpbWUiLCJydW50aW1lU2NoZWR1bGVUaW1lb3V0RXZlbnQiLCJzZXRUaW1lb3V0IiwiX3Jlc3VtZSIsImhhcyIsInJ1bnRpbWVDbGVhclRpbWVvdXRFdmVudCIsImNsZWFyVGltZW91dCIsInJ1bnRpbWVHZXRSYW5kb21EYXRhIiwiZ2V0UmFuZG9tVmFsdWVzIiwic3lzY2FsbEpzRmluYWxpemVSZWYiLCJzeXNjYWxsSnNTdHJpbmdWYWwiLCJzeXNjYWxsSnNWYWx1ZUdldCIsIlJlZmxlY3QiLCJnZXRzcCIsInN5c2NhbGxKc1ZhbHVlU2V0Iiwic3lzY2FsbEpzVmFsdWVEZWxldGUiLCJkZWxldGVQcm9wZXJ0eSIsInN5c2NhbGxKc1ZhbHVlSW5kZXgiLCJzeXNjYWxsSnNWYWx1ZVNldEluZGV4Iiwic3lzY2FsbEpzVmFsdWVDYWxsIiwibSIsInNldFVpbnQ4Iiwic3lzY2FsbEpzVmFsdWVJbnZva2UiLCJzeXNjYWxsSnNWYWx1ZU5ldyIsImNvbnN0cnVjdCIsInN5c2NhbGxKc1ZhbHVlTGVuZ3RoIiwicGFyc2VJbnQiLCJzeXNjYWxsSnNWYWx1ZVByZXBhcmVTdHJpbmciLCJzdHIiLCJlbmNvZGUiLCJzeXNjYWxsSnNWYWx1ZUxvYWRTdHJpbmciLCJzeXNjYWxsSnNWYWx1ZUluc3RhbmNlT2YiLCJzeXNjYWxsSnNDb3B5Qnl0ZXNUb0dvIiwiZHN0IiwiVWludDhDbGFtcGVkQXJyYXkiLCJ0b0NvcHkiLCJzdWJhcnJheSIsInN5c2NhbGxKc0NvcHlCeXRlc1RvSlMiLCJkZWJ1ZyIsIl9ydW4iLCJfdGhpczIiLCJzdHJQdHIiLCJhcmdjIiwiYXJndlB0cnMiLCJ3YXNtTWluRGF0YUFkZHIiLCJJbnN0YW5jZSIsIk5hTiIsImZpbGwiLCJJbmZpbml0eSIsInB0ciIsImJ5dGVzIiwic29ydCIsInJlc3VtZSIs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O0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQSxHQUFHO0FBQ0g7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTs7QUFFQTtBQUNBOztBQUVBO0FBQ0E7QUFDQSxDQUFDOztBQUVEO0FBQ0E7O0FBRUE7QUFDQTtBQUNBLEVBQUU7QUFDRjtBQUNBOztBQUVBOztBQUVBO0FBQ0E7QUFDQSxHQUFHO0FBQ0g7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0EsRUFBRTtBQUNGO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0EsdURBQXVEO0FBQ3ZEOztBQUVBLDZCQUE2QixtQkFBbUI7O0FBRWhEOztBQUVBOztBQUVBO0FBQ0E7Ozs7Ozs7O0FDNVdBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLHdDQUF3QyxXQUFXLEVBQUU7QUFDckQsd0NBQXdDLFdBQVcsRUFBRTs7QUFFckQ7QUFDQTtBQUNBO0FBQ0E7O0FBRUE7QUFDQTs7QUFFQTtBQUNBO0FBQ0E7QUFDQSxHQUFHO0FBQ0g7QUFDQSxzQ0FBc0M7QUFDdEMsR0FBRztBQUNIO0FBQ0EsOERBQThEO0FBQzlEOztBQUVBO0FBQ0E7QUFDQSxFQUFFOztBQUVGO0FBQ0E7QUFDQTs7Ozs7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QUNLLElBQUksQ0FBQ2dSLFlBQVk7Y0FBQTtjQUFBO2dCQUFBLE9BQUF0TSxRQUFBLENBQUE1QyxJQUFBO1lBQUE7VUFBQSxHQUFBd0MsT0FBQTtRQUFBLENBQ3ZCO1FBQUEsU0FBQTBCLElBQUE3QixFQUFBO1VBQUEsT0FBQTBTLElBQUEsQ0FBQTNULEtBQUEsT0FBQUQsU0FBQTtRQUFBO1FBQUEsT0FBQStDLEdBQUE7TUFBQTtJQUFBO01BQUF0RCxHQUFBO01BQUFsRyxLQUFBLEVBRUQsU0FBQWlZLFFBQUEsRUFBVTtRQUNULElBQUksSUFBSSxDQUFDVCxNQUFNLEVBQUU7VUFDaEIsTUFBTSxJQUFJM1UsS0FBSyxDQUFDLCtCQUErQixDQUFDO1FBQ2pEO1FBQ0EsSUFBSSxDQUFDNFQsS0FBSyxDQUFDQyxPQUFPLENBQUN3RSxNQUFNLENBQUMsQ0FBQztRQUMzQixJQUFJLElBQUksQ0FBQzFELE1BQU0sRUFBRTtVQUNoQixJQUFJLENBQUMvQyxtQkFBbUIsQ0FBQyxDQUFDO1FBQzNCO01BQ0Q7SUFBQztNQUFBdk8sR0FBQTtNQUFBbEcsS0FBQSxFQUVELFNBQUFtYixpQkFBaUI3USxFQUFFLEVBQUU7UUFDcEIsSUFBTXhELEVBQUUsR0FBRyxJQUFJO1FBQ2YsT0FBTyxZQUFZO1VBQ2xCLElBQU1zVSxLQUFLLEdBQUc7WUFBRTlRLEVBQUUsRUFBRUEsRUFBRTtZQUFFLFFBQU0sSUFBSTtZQUFFOUQsSUFBSSxFQUFFQztVQUFVLENBQUM7VUFDckRLLEVBQUUsQ0FBQzROLGFBQWEsR0FBRzBHLEtBQUs7VUFDeEJ0VSxFQUFFLENBQUNtUixPQUFPLENBQUMsQ0FBQztVQUNaLE9BQU9tRCxLQUFLLENBQUNyUyxNQUFNO1FBQ3BCLENBQUM7TUFDRjtJQUFDO0lBQUEsT0FBQW9MLE1BQUE7RUFBQSxHQUNEO0FBQ0YsQ0FBQyxFQUFFLENBQUMsQyIsImZpbGUiOiJ3aWRnZXQuanMiLCJzb3VyY2VzQ29udGVudCI6WyIgXHQvLyBUaGUgbW9kdWxlIGNhY2hlXG4gXHR2YXIgaW5zdGFsbGVkTW9kdWxlcyA9IHt9O1xuXG4gXHQvLyBUaGUgcmVxdWlyZSBmdW5jdGlvblxuIFx0ZnVuY3Rpb24gX193ZWJwYWNrX3JlcXVpcmVfXyhtb2R1bGVJZCkge1xuXG4gXHRcdC8vIENoZWNrIGlmIG1vZHVsZSBpcyBpbiBjYWNoZVxuIFx0XHRpZihpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSkge1xuIFx0XHRcdHJldHVybiBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXS5leHBvcnRzO1xuIFx0XHR9XG4gXHRcdC8vIENyZWF0ZSBhIG5ldyBtb2R1bGUgKGFuZCBwdXQgaXQgaW50byB0aGUgY2FjaGUpXG4gXHRcdHZhciBtb2R1bGUgPSBpbnN0YWxsZWRNb2R1bGVzW21vZHVsZUlkXSA9IHtcbiBcdFx0XHRpOiBtb2R1bGVJZCxcbiBcdFx0XHRsOiBmYWxzZSxcbiBcdFx0XHRleHBvcnRzOiB7fVxuIFx0XHR9O1xuXG4gXHRcdC8vIEV4ZWN1dGUgdGhlIG1vZHVsZSBmdW5jdGlvblxuIFx0XHRtb2R1bGVzW21vZHVsZUlkXS5jYWxsKG1vZHVsZS5leHBvcnRzLCBtb2R1bGUsIG1vZHVsZS5leHBvcnRzLCBfX3dlYnBhY2tfcmVxdWlyZV9fKTtcblxuIFx0XHQvLyBGbGFnIHRoZSBtb2R1bGUgYXMgbG9hZGVkXG4gXHRcdG1vZHVsZS5sID0gdHJ1ZTtcblxuIFx0XHQvLyBSZXR1cm4gdGhlIGV4cG9ydHMgb2YgdGhlIG1vZHVsZVxuIFx0XHRyZXR1cm4gbW9kdWxlLmV4cG9ydHM7XG4gXHR9XG5cblxuIFx0Ly8gZXhwb3NlIHRoZSBtb2R1bGVzIG9iamVjdCAoX193ZWJwYWNrX21vZHVsZXNfXylcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubSA9IG1vZHVsZXM7XG5cbiBcdC8vIGV4cG9zZSB0aGUgbW9kdWxlIGNhY2hlXG4gXHRfX3dlYnBhY2tfcmVxdWlyZV9fLmMgPSBpbnN0YWxsZWRNb2R1bGVzO1xuXG4gXHQvLyBkZWZpbmUgZ2V0dGVyIGZ1bmN0aW9uIGZvciBoYXJtb255IGV4cG9ydHNcbiBcdF9fd2VicGFja19yZXF1aXJlX18uZCA9IGZ1bmN0aW9uKGV4cG9ydHMsIG5hbWUsIGdldHRlcikge1xuIFx0XHRpZighX193ZWJwYWNrX3JlcXVpcmVfXy5vKGV4cG9ydHMsIG5hbWUpKSB7XG4gXHRcdFx0T2JqZWN0LmRlZmluZVByb3BlcnR5KGV4cG9ydHMsIG5hbWUsIHtcbiBcdFx0XHRcdGNvbmZpZ3VyYWJsZTogZmFsc2UsXG4gXHRcdFx0XHRlbnVtZXJhYmxlOiB0cnVlLFxuIFx0XHRcdFx0Z2V0OiBnZXR0ZXJcbiBcdFx0XHR9KTtcbiBcdFx0fVxuIFx0fTtcblxuIFx0Ly8gZ2V0RGVmYXVsdEV4cG9ydCBmdW5jdGlvbiBmb3IgY29tcGF0aWJpbGl0eSB3aXRoIG5vbi1oYXJtb255IG1vZHVsZXNcbiBcdF9fd2VicGFja19yZXF1aXJlX18ubiA9IGZ1bmN0aW9uKG1vZHVsZSkge1xuIFx0XHR2YXIgZ2V0dGVyID0gbW9kdWxlICYmIG1vZHVsZS5fX2VzTW9kdWxlID9cbiBcdFx0XHRmdW5jdGlvbiBnZXREZWZhdWx0KCkgeyByZXR1cm4gbW9kdWxlWydkZWZhdWx0J107IH0gOlxuIFx0XHRcdGZ1bmN0aW9uIGdldE1vZHVsZUV4cG9ydHMoKSB7IHJldHVybiBtb2R1bGU7IH07XG4gXHRcdF9fd2VicGFja19yZXF1aXJlX18uZChnZXR0ZXIsICdhJywgZ2V0dGVyKTtcbiBcdFx0cmV0dXJuIGdldHRlcjtcbiBcdH07XG5cbiBcdC8vIE9iamVjdC5wcm90b3R5cGUuaGFzT3duUHJvcGVydHkuY2FsbFxuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5vID0gZnVuY3Rpb24ob2JqZWN0LCBwcm9wZXJ0eSkgeyByZXR1cm4gT2JqZWN0LnByb3RvdHlwZS5oYXNPd25Qcm9wZXJ0eS5jYWxsKG9iamVjdCwgcHJvcGVydHkpOyB9O1xuXG4gXHQvLyBfX3dlYnBhY2tfcHVibGljX3BhdGhfX1xuIFx0X193ZWJwYWNrX3JlcXVpcmVfXy5wID0gXCJcIjtcblxuIFx0Ly8gTG9hZCBlbnRyeSBtb2R1bGUgYW5kIHJldHVybiBleHBvcnRzXG4gXHRyZXR1cm4gX193ZWJwYWNrX3JlcXVpcmVfXyhfX3dlYnBhY2tfcmVxdWlyZV9fLnMgPSAwKTtcblxuXG5cbi8vIFdFQlBBQ0sgRk9PVEVSIC8vXG4vLyB3ZWJwYWNrL2Jvb3RzdHJhcCBjMDk0ZmI5OTEyODIwMDUyNzE2MCIsImltcG9ydCB7IHBpbmcgfSBmcm9tICcuL3NlcnZpY2VzJ1xuLy8gaW1wb3J0IHsgc2hvdyB9IGZyb20gJy4vdmlld3MvbWVzc2FnZSdcbmltcG9ydCB7IGFkZEJ1dHRvblRvVGV4dGFyZWEgfSBmcm9tICcuL3ZpZXdzL2xiJztcblxuY29uc3Qgc3VwcG9ydGVkQVBJID0gWydpbml0JywnbGFtYWJ1dHRvbiddOyAvLyBlbmxpc3QgYWxsIG1ldGhvZHMgc3VwcG9ydGVkIGJ5IEFQSSAoZS5nLiBgbXcoJ2V2ZW50JywgJ3VzZXItbG9naW4nKTtgKVxubGV0IGdvOyBcbi8qKlxuICAgIFRoZSBtYWluIGVudHJ5IG9mIHRoZSBhcHBsaWNhdGlvblxuICAgICovXG5mdW5jdGlvbiBhcHAod2luZG93KSB7XG4gICAgY29uc29sZS5sb2coJ0pTLVdpZGdldCBzdGFydGluZycpO1xuXG4gICAgLy8gc2V0IGRlZmF1bHQgY29uZmlndXJhdGlvbnNcbiAgICBsZXQgY29uZmlndXJhdGlvbnMgPSB7XG4gICAgICAgIHNvbWVEZWZhdWx0Q29uZmlndXJhdGlvbjogZmFsc2VcbiAgICB9O1xuXG4gICAgLy8gYWxsIG1ldGhvZHMgdGhhdCB3ZXJlIGNhbGxlZCB0aWxsIG5vdyBhbmQgc3RvcmVkIGluIHF1ZXVlXG4gICAgLy8gbmVlZHMgdG8gYmUgY2FsbGVkIG5vdyBcbiAgICBsZXQgZ2xvYmFsT2JqZWN0ID0gd2luZG93W3dpbmRvd1snSlMtV2lkZ2V0J11dO1xuICAgIGxldCBxdWV1ZSA9IGdsb2JhbE9iamVjdC5xO1xuICAgIGlmIChxdWV1ZSkge1xuICAgICAgICBmb3IgKHZhciBpID0gMDsgaSA8IHF1ZXVlLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgICAgICBpZiAocXVldWVbaV1bMF0udG9Mb3dlckNhc2UoKSA9PSAnaW5pdCcpIHtcbiAgICAgICAgICAgICAgICBjb25maWd1cmF0aW9ucyA9IGV4dGVuZE9iamVjdChjb25maWd1cmF0aW9ucywgcXVldWVbaV1bMV0pO1xuICAgICAgICAgICAgICAgIGNvbnNvbGUubG9nKCdKUy1XaWRnZXQgc3RhcnRlZCcsIGNvbmZpZ3VyYXRpb25zKTtcbiAgICAgICAgICAgIH1cbiAgICAgICAgICAgIGVsc2VcbiAgICAgICAgICAgICAgICBhcGlIYW5kbGVyKHF1ZXVlW2ldWzBdLCBxdWV1ZVtpXVsxXSk7XG4gICAgICAgIH1cbiAgICB9XG5cbiAgICAvLyBvdmVycmlkZSB0ZW1wb3JhcnkgKHVudGlsIHRoZSBhcHAgbG9hZGVkKSBoYW5kbGVyXG4gICAgLy8gZm9yIHdpZGdldCdzIEFQSSBjYWxsc1xuICAgIGdsb2JhbE9iamVjdCA9IGFwaUhhbmRsZXI7XG4gICAgZ2xvYmFsT2JqZWN0LmNvbmZpZ3VyYXRpb25zID0gY29uZmlndXJhdGlvbnM7XG59XG5cbi8qKlxuICAgIE1ldGhvZCB0aGF0IGhhbmRsZXMgYWxsIEFQSSBjYWxsc1xuICAgICovXG5hc3luYyBmdW5jdGlvbiBhcGlIYW5kbGVyKGFwaSwgcGFyYW1zKSB7XG4gICAgaWYgKCFhcGkpIHRocm93IEVycm9yKCdBUEkgbWV0aG9kIHJlcXVpcmVkJyk7XG4gICAgYXBpID0gYXBpLnRvTG93ZXJDYXNlKCk7XG5cbiAgICBpZiAoc3VwcG9ydGVkQVBJLmluZGV4T2YoYXBpKSA9PT0gLTEpIHRocm93IEVycm9yKGBNZXRob2QgJHthcGl9IGlzIG5vdCBzdXBwb3J0ZWRgKTtcblxuICAgIGNvbnNvbGUubG9nKGBIYW5kbGluZyBBUEkgY2FsbCAke2FwaX1gLCBwYXJhbXMpO1xuXG4gICAgc3dpdGNoIChhcGkpIHtcbiAgICAgICAgLy8gVE9ETzogYWRkIEFQSSBpbXBsZW1lbnRhdGlvblxuICAgICAgICBjYXNlICdsYW1hYnV0dG9uJzpcbiAgICAgICAgICAgIGdvID0gYXdhaXQgbG9hZFdhc20oKTsgLy8gRW5zdXJlIFdhc20gaXMgbG9hZGVkIGFuZCBHbyBpbnN0YW5jZSBpcyBjcmVhdGVkXG4gICAgICAgICAgICBhZGRCdXR0b25Ub1RleHRhcmVhKHBhcmFtcyxnbyk7XG4gICAgICAgICAgICBicmVhaztcbiAgICAgICAgZGVmYXVsdDpcbiAgICAgICAgICAgIGNvbnNvbGUud2FybihgTm8gaGFuZGxlciBkZWZpbmVkIGZvciAke2FwaX1gKTtcbiAgICB9XG59XG5cbmZ1bmN0aW9uIGV4dGVuZE9iamVjdChhLCBiKSB7XG4gICAgZm9yICh2YXIga2V5IGluIGIpXG4gICAgICAgIGlmIChiLmhhc093blByb3BlcnR5KGtleSkpXG4gICAgICAgICAgICBhW2tleV0gPSBiW2tleV07XG4gICAgcmV0dXJuIGE7XG59XG5cbmFwcCh3aW5kb3cpO1xuXG5hc3luYyBmdW5jdGlvbiBsb2FkV2FzbSgpIHtcbiAgICB0cnkge1xuICAgICAgICBjb25zdCByZXNwb25zZSA9IGF3YWl0IGZldGNoKFwibWFpbi53YXNtXCIpO1xuICAgICAgICBjb25zdCBidWZmZXIgPSBhd2FpdCByZXNwb25zZS5hcnJheUJ1ZmZlcigpO1xuICAgICAgICBjb25zdCBnbyA9IG5ldyBHbygpO1xuICAgICAgICBjb25zdCByZXN1bHQgPSBhd2FpdCBXZWJBc3NlbWJseS5pbnN0YW50aWF0ZShidWZmZXIsIGdvLmltcG9ydE9iamVjdCk7XG4gICAgICAgIGdvLnJ1bihyZXN1bHQuaW5zdGFuY2UpO1xuICAgICAgICBjb25zb2xlLmxvZyhcIldhc20gbG9hZGVkIGFuZCBHbyBpbnN0YW5jZSBpbml0aWFsaXplZFwiKTtcbiAgICAgICAgcmV0dXJuIGdvO1xuICAgIH0gY2F0Y2ggKGVycm9yKSB7XG4gICAgICAgIGNvbnNvbGUuZXJyb3IoXCJFcnJvciBsb2FkaW5nIFdhc206XCIsIGVycm9yKTtcbiAgICAgICAgcmV0dXJuIG51bGw7XG4gICAgfVxufVxuXG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL21haW4uanMiLCJcbmV4cG9ydCBmdW5jdGlvbiBwaW5nKCkge1xuICAgIHJldHVybiAncG9uZyc7XG59XG5cblxuLy8gV0VCUEFDSyBGT09URVIgLy9cbi8vIC4vc3JjL3NlcnZpY2VzLmpzIiwiaW1wb3J0ICcuL2xiLmNzcyc7XG5pbXBvcnQgJy4vd2FzbV9leGVjJ1xuXG5leHBvcnQgZnVuY3Rpb24gYWRkQnV0dG9uVG9UZXh0YXJlYShjbGFzc05hbWUsZ28pIHtcbiAgICBpZiAoIWdvKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKFwid2FzbSBmaWxlIG5vdCBpbml0aWFsaXplZFwiKVxuICAgICAgICAgICAgcmV0dXJuOyBcbiAgICAgICAgfVxuXG4gICAgY29uc3QgbDJjb2RlRWxlbWVudHMgPSBkb2N1bWVudC5xdWVyeVNlbGVjdG9yQWxsKGAuJHtjbGFzc05hbWV9YCk7XG5cbiAgICBsMmNvZGVFbGVtZW50cy5mb3JFYWNoKChkaXZFbGVtZW50KSA9PiB7XG4gICAgICAgIGNvbnN0IGwyY29udGVudCA9IGRpdkVsZW1lbnQudGV4dENvbnRlbnQ7XG4gICAgICAgIGRpdkVsZW1lbnQudGV4dENvbnRlbnQgPSAnJztcblxuICAgICAgICBjb25zdCBlbGVtZW50VHlwZSA9IGRpdkVsZW1lbnQudGFnTmFtZTsgLy8gR2V0IHRoZSB0YWcgbmFtZSBvZiB0aGUgZGl2RWxlbWVudFxuICAgICAgICBjb25zb2xlLmxvZygnRWxlbWVudCB0eXBlOicsIGVsZW1lbnRUeXBlKVxuXG4gICAgICAgIGNvbnN0IHRleHRhcmVhRWxlbWVudCA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ3RleHRhcmVhJyk7XG4gICAgICAgIHRleHRhcmVhRWxlbWVudC5pZCA9ICdteVRleHRhcmVhJztcbiAgICAgICAgdGV4dGFyZWFFbGVtZW50Lm5hbWUgPSAnSGFzaCc7XG4gICAgICAgIHRleHRhcmVhRWxlbWVudC5zdHlsZS5tYXhXaWR0aCA9ICcxMDAlJzsgLy8gU2V0IGEgbWF4aW11bSB3aWR0aCB0byBtYWtlIGl0IHJlc3BvbnNpdmVcbiAgICAgICAgdGV4dGFyZWFFbGVtZW50LnN0eWxlLndpZHRoID0gJzUwMHB4JztcbiAgICAgICAgdGV4dGFyZWFFbGVtZW50LnN0eWxlLmhlaWdodCA9ICcyMDBweCc7XG4gICAgICAgIHRleHRhcmVhRWxlbWVudC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSAnI2YwZjBmMCc7IFxuICAgICAgICB0ZXh0YXJlYUVsZW1lbnQuc3R5bGUucGFkZGluZyA9ICcxMHB4JztcbiAgICAgICAgdGV4dGFyZWFFbGVtZW50LnRleHRDb250ZW50ID0gbDJjb250ZW50O1xuICAgICAgICBkaXZFbGVtZW50LmFwcGVuZENoaWxkKHRleHRhcmVhRWxlbWVudCk7XG5cbiAgICAgICAgXG4gICAgICAgIGNvbnN0IGJ1dHRvbkltYWdlID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnaW1nJyk7XG4gICAgICAgIGJ1dHRvbkltYWdlLnNyYyA9ICdhcnJvd2wyLnBuZyc7XG4gICAgICAgIGJ1dHRvbkltYWdlLnN0eWxlLndpZHRoID0gJzMwcHgnOyBcbiAgICAgICAgY29uc3QgaW5kZXggPSBNYXRoLnJhbmRvbSgpLnRvU3RyaW5nKDM2KS5zdWJzdHIoMiwgMTApO1xuICAgICAgICBjb25zdCBidXR0b25FbGVtZW50ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnYnV0dG9uJyk7XG4gICAgICAgIGJ1dHRvbkVsZW1lbnQuYXBwZW5kQ2hpbGQoYnV0dG9uSW1hZ2UpO1xuICAgICAgICBidXR0b25FbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ2NsaWNrJywgKCkgPT4gYWRkVGV4dFRvRGl2KGRpdkVsZW1lbnQsaW5kZXgsZ28pKTtcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5hbGwgPSAndW5zZXQnO1xuICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLnBvc2l0aW9uID0gJ3JlbGF0aXZlJztcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5ib3R0b20gPSAnMTUwcHgnO1xuICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLnJpZ2h0ID0gJzUwcHgnO1xuICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLnRyYW5zaXRpb24gPSAnYmFja2dyb3VuZC1jb2xvciAwLjNzJztcbiAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5jdXJzb3IgPSAncG9pbnRlcic7XG4gICAgICAgIGNvbnN0IGhpZ2hsaWdodENvbG9yID0gJyMwMEExRTEnO1xuICAgICAgICBidXR0b25FbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNlZW50ZXInLCAoKSA9PiB7XG4gICAgICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLmJhY2tncm91bmRDb2xvciA9IGhpZ2hsaWdodENvbG9yOyAvLyBDaGFuZ2UgdGhlIGJhY2tncm91bmQgY29sb3Igb24gaG92ZXJcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBcbiAgICAgICAgICBidXR0b25FbGVtZW50LmFkZEV2ZW50TGlzdGVuZXIoJ21vdXNlbGVhdmUnLCAoKSA9PiB7XG4gICAgICAgICAgICBidXR0b25FbGVtZW50LnN0eWxlLmJhY2tncm91bmRDb2xvciA9ICcnOyAvLyBSZXNldCB0aGUgYmFja2dyb3VuZCBjb2xvciBvbiBtb3VzZSBsZWF2ZVxuICAgICAgICAgIH0pO1xuICAgICAgICAgIFxuICAgICAgICAgIGJ1dHRvbkVsZW1lbnQuYWRkRXZlbnRMaXN0ZW5lcignbW91c2Vkb3duJywgKCkgPT4ge1xuICAgICAgICAgICAgYnV0dG9uRWxlbWVudC5zdHlsZS5iYWNrZ3JvdW5kQ29sb3IgPSBoaWdobGlnaHRDb2xvcjsgLy8gQ2hhbmdlIHRoZSBiYWNrZ3JvdW5kIGNvbG9yIHdoZW4gY2xpY2tlZFxuICAgICAgICAgIH0pO1xuICAgICAgICAgIFxuICAgICAgICAgIGJ1dHRvbkVsZW1lbnQuYWRkRXZlbnRMaXN0ZW5lcignbW91c2V1cCcsICgpID0+IHtcbiAgICAgICAgICAgIGJ1dHRvbkVsZW1lbnQuc3R5bGUuYmFja2dyb3VuZENvbG9yID0gaGlnaGxpZ2h0Q29sb3I7IC8vIFJlc3RvcmUgdGhlIGhvdmVyIGNvbG9yIG9uIHJlbGVhc2VcbiAgICAgICAgICB9KTtcbiAgICAgICAgICBcblxuICAgICAgICBkaXZFbGVtZW50LmFwcGVuZENoaWxkKGJ1dHRvbkVsZW1lbnQpO1xuXG4gICAgICAgIGNvbnN0IGxvZ29JbWFnZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2ltZycpO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUucG9zaXRpb24gPSAncmVsYXRpdmUnO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUuYm90dG9tID0gJzEwcHgnO1xuICAgICAgICBsb2dvSW1hZ2Uuc3R5bGUucmlnaHQgPSAnODBweCc7XG4gICAgICAgIGxvZ29JbWFnZS5zcmMgPSAnaGV4bW9zX2xvZ28uanBlZyc7XG4gICAgICAgIGxvZ29JbWFnZS5zdHlsZS53aWR0aCA9ICczMHB4JztcbiAgICAgICAgbG9nb0ltYWdlLnN0eWxlLmJvcmRlclJhZGl1cyA9ICc1MCUnOyBcblxuICAgICAgICBsb2dvSW1hZ2UuYWRkRXZlbnRMaXN0ZW5lcignY2xpY2snLCBmdW5jdGlvbigpIHtcbiAgICAgICAgICAgIHdpbmRvdy5sb2NhdGlvbi5ocmVmID0gJ2h0dHBzOi8vaGV4bW9zLmNvbSc7XG4gICAgICAgICAgfSk7XG4gICAgICAgICAgXG4gICAgICAgIGRpdkVsZW1lbnQuYXBwZW5kQ2hpbGQobG9nb0ltYWdlKVxuXG4gICAgfSk7XG59XG5cbmFzeW5jIGZ1bmN0aW9uIGFkZFRleHRUb0RpdihkaXZFbGVtZW50LGluZGV4LGdvKSB7XG4gICAgY29uc3QgdGV4dEFyZWEgPSBkaXZFbGVtZW50LnF1ZXJ5U2VsZWN0b3IoJ3RleHRhcmVhJyk7XG4gICAgY29uc3Qgb3V0cHV0SWQgPSBgb3V0cHV0RmllbGRfJHtpbmRleH1gO1xuICAgIGNvbnN0IGV4aXN0aW5nT3V0cHV0RGl2ID0gZG9jdW1lbnQuZ2V0RWxlbWVudEJ5SWQob3V0cHV0SWQpO1xuXG4gICAgaWYgKHRleHRBcmVhKSB7XG4gICAgICAgIGNvbnNvbGUubG9nKFwiVGV4dCBpbnNpZGUgdGhlIHRleHQgYXJlYTpcIiwgdGV4dEFyZWEudmFsdWUpO1xuICAgICAgICBpZiAoIWV4aXN0aW5nT3V0cHV0RGl2KSB7XG4gICAgICAgICAgICBjb25zb2xlLmxvZyhcIlRleHQgaW5zaWRlIHRoZSB0ZXh0IGFyZWE6XCIsIHRleHRBcmVhLnZhbHVlKTtcbiAgICAgICAgICAgIGNvbnN0IG91dHB1dERpdiA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoJ2RpdicpO1xuICAgICAgICAgICAgb3V0cHV0RGl2LmNsYXNzTmFtZSA9ICdvdXRwdXRGaWVsZCc7XG4gICAgICAgICAgICBvdXRwdXREaXYuaWQgPSBgb3V0cHV0RmllbGRfJHtpbmRleH1gO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmZvbnRTaXplID0gJzE2cHgnO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLm1hcmdpblRvcCA9ICcxMHB4JztcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS53aWR0aCA9ICc1MDBweCc7IFxuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLm1pbkhlaWdodCA9ICcyMDBweCc7IFxuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmNvbG9yID0gJ3doaXRlJztcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS5oZWlnaHQgPSAnNDAwcHgnOyBcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS5vdmVyZmxvd1kgPSAnYXV0byc7IFxuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmJhY2tncm91bmRDb2xvciA9ICcjMzMzJztcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS5ib3JkZXIgPSAnMXB4IHNvbGlkICM1NTUnO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmJvcmRlclJhZGl1cyA9ICc1cHgnOyBcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS5wYWRkaW5nID0gJzIwcHgnOyBcbiAgICAgICAgICAgIG91dHB1dERpdi5zdHlsZS5mb250RmFtaWx5ID0gJ0FyaWFsLCBzYW5zLXNlcmlmJzsgXG4gICAgICAgICAgICBvdXRwdXREaXYuc3R5bGUub3ZlcmZsb3dZID0gJ2F1dG8nO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLndoaXRlU3BhY2UgPSAncHJlLXdyYXAnO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLmJveFNoYWRvdyA9ICcwIDAgMTBweCByZ2JhKDAsIDAsIDAsIDAuMiknO1xuICAgICAgICAgICAgb3V0cHV0RGl2LnN0eWxlLnRyYW5zaXRpb24gPSAnYmFja2dyb3VuZC1jb2xvciAwLjNzJztcbiAgICAgICAgXG4gICAgICAgICAgICBkaXZFbGVtZW50LmFwcGVuZENoaWxkKG91dHB1dERpdik7XG4gICAgICAgICAgICAvLyBtYWtlTGFtYVJlcXVlc3QodGV4dEFyZWEudmFsdWUsIGRpdkVsZW1lbnQsIGluZGV4KTtcbiAgICAgICAgfVxuICAgICAgICBtYWtlTGFtYVJlcXVlc3QodGV4dEFyZWEudmFsdWUsIGRpdkVsZW1lbnQsIGluZGV4LGdvKVxuICAgIH1cbiAgICBlbHNle1xuICAgICAgICBjb25zb2xlLmxvZyhcIk5vIHRleHQgYXJlYSBmb3VuZC4gTG9nZ2luZyBkYXRhIGZyb20gZGl2RWxlbWVudDpcIiwgZGl2RWxlbWVudC50ZXh0Q29udGVudCk7XG4gICAgICAgIG1ha2VMYW1hUmVxdWVzdChkaXZFbGVtZW50LnRleHRDb250ZW50LGRpdkVsZW1lbnQsIGluZGV4KVxuICAgIH1cbn1cblxuYXN5bmMgZnVuY3Rpb24gbWFrZUxhbWFSZXF1ZXN0KGNvbW1hbmQsIGRpdkVsZW1lbnQsIGluZGV4LGdvKSB7XG4gICAgdHJ5IHtcbiAgICAgICAgaWYgKCFnbykge1xuICAgICAgICAgICAgY29uc29sZS5lcnJvcihcIldhc20gaXMgbm90IGxvYWRlZC4gQ2Fubm90IG1ha2UgYSByZXF1ZXN0LlwiKTtcbiAgICAgICAgICAgIHJldHVybjtcbiAgICAgICAgfVxuICAgICAgICBjb25zb2xlLmxvZyhcImdvIG9iamVjdCA6XCIsZ28pXG4gICAgICAgIGNvbnN0IHJlc3BvbnNlID0gYXdhaXQgZ29XZWJSZXF1ZXN0RnVuYyhjb21tYW5kKTtcbiAgICAgICAgY29uc29sZS5sb2coXCJyZXNwb25zZSBEYXRhOlwiLHJlc3BvbnNlKVxuICAgICAgICBjb25zdCBqc29uT2JqZWN0ID0gSlNPTi5wYXJzZShyZXNwb25zZSk7XG4gICAgICAgIGNvbnN0IGZvcm1hdHRlZEJvZHkgPSBKU09OLnBhcnNlKGpzb25PYmplY3QuQm9keSk7XG4gICAgICAgIGNvbnN0IHJlc3BTdGF0dXNDb2RlID0gIEpTT04ucGFyc2UoanNvbk9iamVjdC5TdGF0dXNDb2RlKTtcbiAgICAgICAgdmFyIGZvcm1hdHRlZEpzb24gPSBKU09OLnN0cmluZ2lmeShmb3JtYXR0ZWRCb2R5LCBudWxsLCAyKTtcbiAgICAgICAgLy8gdmFyIGZvcm1hdHRlZEpzb24gPSBKU09OLnN0cmluZ2lmeShqc29uT2JqZWN0LCBudWxsLCAyKTtcbiAgICAgICAgY29uc29sZS5sb2coXCJmb3JtYXR0ZWRKc29uOlwiLGZvcm1hdHRlZEpzb24pICAgIFxuICAgICAgICBjb25zdCBvdXRwdXRGaWVsZCA9IGRpdkVsZW1lbnQucXVlcnlTZWxlY3RvcihgI291dHB1dEZpZWxkXyR7aW5kZXh9YCk7XG4gICAgICAgIGNvbnNvbGUubG9nKG91dHB1dEZpZWxkKVxuICAgICAgICBvdXRwdXRGaWVsZC50ZXh0Q29udGVudCA9IGZvcm1hdHRlZEpzb247IFxuICAgICAgICBjb25zb2xlLmxvZyhcIm91dHB1dEZpZWxkLnRleHRDb250ZW50OlwiLG91dHB1dEZpZWxkLnRleHRDb250ZW50KVxuICAgICAgICBjb25zdCBzdGF0dXNDb2RlRGl2ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudCgnZGl2Jyk7XG4gICAgICAgIHN0YXR1c0NvZGVEaXYudGV4dENvbnRlbnQgPSBgU3RhdHVzIENvZGU6ICR7cmVzcFN0YXR1c0NvZGV9YDtcbiAgICAgICAgaWYgKHJlc3BTdGF0dXNDb2RlID4gMjAwKSB7XG4gICAgICAgICAgICBzdGF0dXNDb2RlRGl2LnN0eWxlLmNvbG9yID0gJ3JlZCc7IC8vIFNldCB0ZXh0IGNvbG9yIHRvIHJlZCBmb3Igc3RhdHVzIGNvZGVzID4gMjAwXG4gICAgICAgIH0gZWxzZSB7XG4gICAgICAgICAgICBzdGF0dXNDb2RlRGl2LnN0eWxlLmNvbG9yID0gJ2dyZWVuJzsgLy8gU2V0IHRleHQgY29sb3IgdG8gZ3JlZW4gZm9yIHN0YXR1cyBjb2RlcyA8PSAyMDBcbiAgICAgICAgfVxuICAgICAgICAvLyBzdGF0dXNDb2RlRGl2LnN0eWxlLmNvbG9yID0gJ2dyZWVuJzsgLy8gU2V0IHRleHQgY29sb3IgdG8gZ3JlZW5cbiAgICAgICAgc3RhdHVzQ29kZURpdi5zdHlsZS5mb250V2VpZ2h0ID0gJ2JvbGQnOyAvLyBNYWtlIHRoZSB0ZXh0IGJvbGRcbiAgICAgICAgLy8gZGl2RWxlbWVudC5hcHBlbmRDaGlsZChzdGF0dXNDb2RlRGl2KTtcbiAgICAgICAgZGl2RWxlbWVudC5pbnNlcnRCZWZvcmUoc3RhdHVzQ29kZURpdiwgb3V0cHV0RmllbGQpO1xuXG4gICAgfSBjYXRjaCAoZXJyb3IpIHtcbiAgICAgICAgY29uc29sZS5lcnJvcihcIkVycm9yIG1ha2luZyBMYW1hIHJlcXVlc3Q6XCIsIGVycm9yKTtcbiAgICB9XG59XG5cblxuXG4vLyBXRUJQQUNLIEZPT1RFUiAvL1xuLy8gLi9zcmMvdmlld3MvbGIuanMiLCIvLyBzdHlsZS1sb2FkZXI6IEFkZHMgc29tZSBjc3MgdG8gdGhlIERPTSBieSBhZGRpbmcgYSA8c3R5bGU+IHRhZ1xuXG4vLyBsb2FkIHRoZSBzdHlsZXNcbnZhciBjb250ZW50ID0gcmVxdWlyZShcIiEhLi4vLi4vbm9kZV9tb2R1bGVzL2Nzcy1sb2FkZXIvZGlzdC9janMuanMhLi9sYi5jc3NcIik7XG5pZih0eXBlb2YgY29udGVudCA9PT0gJ3N0cmluZycpIGNvbnRlbnQgPSBbW21vZHVsZS5pZCwgY29udGVudCwgJyddXTtcbi8vIFByZXBhcmUgY3NzVHJhbnNmb3JtYXRpb25cbnZhciB0cmFuc2Zvcm07XG5cbnZhciBvcHRpb25zID0ge1wiaG1yXCI6dHJ1ZX1cbm9wdGlvbnMudHJhbnNmb3JtID0gdHJhbnNmb3JtXG4vLyBhZGQgdGhlIHN0eWxlcyB0byB0aGUgRE9NXG52YXIgdXBkYXRlID0gcmVxdWlyZShcIiEuLi8uLi9ub2RlX21vZHVsZXMvc3R5bGUtbG9hZGVyL2xpYi9hZGRTdHlsZXMuanNcIikoY29udGVudCwgb3B0aW9ucyk7XG5pZihjb250ZW50LmxvY2FscykgbW9kdWxlLmV4cG9ydHMgPSBjb250ZW50LmxvY2Fscztcbi8vIEhvdCBNb2R1bGUgUmVwbGFjZW1lbnRcbmlmKG1vZHVsZS5ob3QpIHtcblx0Ly8gV2hlbiB0aGUgc3R5bGVzIGNoYW5nZSwgdXBkYXRlIHRoZSA8c3R5bGU+IHRhZ3Ncblx0aWYoIWNvbnRlbnQubG9jYWxzKSB7XG5cdFx0bW9kdWxlLmhvdC5hY2NlcHQoXCIhIS4uLy4uL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvY2pzLmpzIS4vbGIuY3NzXCIsIGZ1bmN0aW9uKCkge1xuXHRcdFx0dmFyIG5ld0NvbnRlbnQgPSByZXF1aXJlKFwiISEuLi8uLi9ub2RlX21vZHVsZXMvY3NzLWxvYWRlci9kaXN0L2Nqcy5qcyEuL2xiLmNzc1wiKTtcblx0XHRcdGlmKHR5cGVvZiBuZXdDb250ZW50ID09PSAnc3RyaW5nJykgbmV3Q29udGVudCA9IFtbbW9kdWxlLmlkLCBuZXdDb250ZW50LCAnJ11dO1xuXHRcdFx0dXBkYXRlKG5ld0NvbnRlbnQpO1xuXHRcdH0pO1xuXHR9XG5cdC8vIFdoZW4gdGhlIG1vZHVsZSBpcyBkaXNwb3NlZCwgcmVtb3ZlIHRoZSA8c3R5bGU+IHRhZ3Ncblx0bW9kdWxlLmhvdC5kaXNwb3NlKGZ1bmN0aW9uKCkgeyB1cGRhdGUoKTsgfSk7XG59XG5cblxuLy8vLy8vLy8vLy8vLy8vLy8vXG4vLyBXRUJQQUNLIEZPT1RFUlxuLy8gLi9zcmMvdmlld3MvbGIuY3NzXG4vLyBtb2R1bGUgaWQgPSAzXG4vLyBtb2R1bGUgY2h1bmtzID0gMCIsIi8vIEltcG9ydHNcbnZhciBfX19DU1NfTE9BREVSX0FQSV9JTVBPUlRfX18gPSByZXF1aXJlKFwiLi4vLi4vbm9kZV9tb2R1bGVzL2Nzcy1sb2FkZXIvZGlzdC9ydW50aW1lL2FwaS5qc1wiKTtcbmV4cG9ydHMgPSBfX19DU1NfTE9BREVSX0FQSV9JTVBPUlRfX18oZmFsc2UpO1xuLy8gTW9kdWxlXG5leHBvcnRzLnB1c2goW21vZHVsZS5pZCwgXCIvKiBsYW1hYnV0dG9uLmNzcyAqL1xcblxcbi8qIFN0eWxlIGZvciB0aGUgYnV0dG9uIGVsZW1lbnQgKi9cXG4uanMtd2lkZ2V0LWxhbWFidXR0b24ge1xcbiAgICBiYWNrZ3JvdW5kLWNvbG9yOiBncmVlbjsgLyogU2V0IHRoZSBiYWNrZ3JvdW5kIGNvbG9yIHRvIGdyZWVuICovXFxuICAgIGNvbG9yOiB3aGl0ZTsgLyogU2V0IHRoZSB0ZXh0IGNvbG9yIHRvIHdoaXRlICovXFxuICAgIGJvcmRlcjogbm9uZTsgLyogUmVtb3ZlIGJvcmRlcnMgKi9cXG4gICAgcGFkZGluZzogNXB4IDEwcHg7IC8qIEFkanVzdCBwYWRkaW5nIGZvciBhIHNtYWxsZXIgYnV0dG9uICovXFxuICAgIGN1cnNvcjogcG9pbnRlcjsgLyogQWRkIGEgcG9pbnRlciBjdXJzb3Igb24gaG92ZXIgKi9cXG4gICAgcG9zaXRpb246IGFic29sdXRlO1xcbn1cXG5cXG4vKiBIb3ZlciBlZmZlY3Qgd2hlbiB0aGUgYnV0dG9uIGlzIGhvdmVyZWQgb3ZlciAqL1xcbi5qcy13aWRnZXQtbGFtYWJ1dHRvbjpob3ZlciB7XFxuICAgIGJhY2tncm91bmQtY29sb3I6IGRhcmtncmVlbjsgLyogRGFya2VuIHRoZSBiYWNrZ3JvdW5kIGNvbG9yIG9uIGhvdmVyICovXFxufVxcblwiLCBcIlwiXSk7XG4vLyBFeHBvcnRzXG5tb2R1bGUuZXhwb3J0cyA9IGV4cG9ydHM7XG5cblxuXG4vLy8vLy8vLy8vLy8vLy8vLy9cbi8vIFdFQlBBQ0sgRk9PVEVSXG4vLyAuL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvY2pzLmpzIS4vc3JjL3ZpZXdzL2xiLmNzc1xuLy8gbW9kdWxlIGlkID0gNFxuLy8gbW9kdWxlIGNodW5rcyA9IDAiLCJcInVzZSBzdHJpY3RcIjtcblxuLypcbiAgTUlUIExpY2Vuc2UgaHR0cDovL3d3dy5vcGVuc291cmNlLm9yZy9saWNlbnNlcy9taXQtbGljZW5zZS5waHBcbiAgQXV0aG9yIFRvYmlhcyBLb3BwZXJzIEBzb2tyYVxuKi9cbi8vIGNzcyBiYXNlIGNvZGUsIGluamVjdGVkIGJ5IHRoZSBjc3MtbG9hZGVyXG4vLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgZnVuYy1uYW1lc1xubW9kdWxlLmV4cG9ydHMgPSBmdW5jdGlvbiAodXNlU291cmNlTWFwKSB7XG4gIHZhciBsaXN0ID0gW107IC8vIHJldHVybiB0aGUgbGlzdCBvZiBtb2R1bGVzIGFzIGNzcyBzdHJpbmdcblxuICBsaXN0LnRvU3RyaW5nID0gZnVuY3Rpb24gdG9TdHJpbmcoKSB7XG4gICAgcmV0dXJuIHRoaXMubWFwKGZ1bmN0aW9uIChpdGVtKSB7XG4gICAgICB2YXIgY29udGVudCA9IGNzc1dpdGhNYXBwaW5nVG9TdHJpbmcoaXRlbSwgdXNlU291cmNlTWFwKTtcblxuICAgICAgaWYgKGl0ZW1bMl0pIHtcbiAgICAgICAgcmV0dXJuIFwiQG1lZGlhIFwiLmNvbmNhdChpdGVtWzJdLCBcIiB7XCIpLmNvbmNhdChjb250ZW50LCBcIn1cIik7XG4gICAgICB9XG5cbiAgICAgIHJldHVybiBjb250ZW50O1xuICAgIH0pLmpvaW4oJycpO1xuICB9OyAvLyBpbXBvcnQgYSBsaXN0IG9mIG1vZHVsZXMgaW50byB0aGUgbGlzdFxuICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgZnVuYy1uYW1lc1xuXG5cbiAgbGlzdC5pID0gZnVuY3Rpb24gKG1vZHVsZXMsIG1lZGlhUXVlcnksIGRlZHVwZSkge1xuICAgIGlmICh0eXBlb2YgbW9kdWxlcyA9PT0gJ3N0cmluZycpIHtcbiAgICAgIC8vIGVzbGludC1kaXNhYmxlLW5leHQtbGluZSBuby1wYXJhbS1yZWFzc2lnblxuICAgICAgbW9kdWxlcyA9IFtbbnVsbCwgbW9kdWxlcywgJyddXTtcbiAgICB9XG5cbiAgICB2YXIgYWxyZWFkeUltcG9ydGVkTW9kdWxlcyA9IHt9O1xuXG4gICAgaWYgKGRlZHVwZSkge1xuICAgICAgZm9yICh2YXIgaSA9IDA7IGkgPCB0aGlzLmxlbmd0aDsgaSsrKSB7XG4gICAgICAgIC8vIGVzbGludC1kaXNhYmxlLW5leHQtbGluZSBwcmVmZXItZGVzdHJ1Y3R1cmluZ1xuICAgICAgICB2YXIgaWQgPSB0aGlzW2ldWzBdO1xuXG4gICAgICAgIGlmIChpZCAhPSBudWxsKSB7XG4gICAgICAgICAgYWxyZWFkeUltcG9ydGVkTW9kdWxlc1tpZF0gPSB0cnVlO1xuICAgICAgICB9XG4gICAgICB9XG4gICAgfVxuXG4gICAgZm9yICh2YXIgX2kgPSAwOyBfaSA8IG1vZHVsZXMubGVuZ3RoOyBfaSsrKSB7XG4gICAgICB2YXIgaXRlbSA9IFtdLmNvbmNhdChtb2R1bGVzW19pXSk7XG5cbiAgICAgIGlmIChkZWR1cGUgJiYgYWxyZWFkeUltcG9ydGVkTW9kdWxlc1tpdGVtWzBdXSkge1xuICAgICAgICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgbm8tY29udGludWVcbiAgICAgICAgY29udGludWU7XG4gICAgICB9XG5cbiAgICAgIGlmIChtZWRpYVF1ZXJ5KSB7XG4gICAgICAgIGlmICghaXRlbVsyXSkge1xuICAgICAgICAgIGl0ZW1bMl0gPSBtZWRpYVF1ZXJ5O1xuICAgICAgICB9IGVsc2Uge1xuICAgICAgICAgIGl0ZW1bMl0gPSBcIlwiLmNvbmNhdChtZWRpYVF1ZXJ5LCBcIiBhbmQgXCIpLmNvbmNhdChpdGVtWzJdKTtcbiAgICAgICAgfVxuICAgICAgfVxuXG4gICAgICBsaXN0LnB1c2goaXRlbSk7XG4gICAgfVxuICB9O1xuXG4gIHJldHVybiBsaXN0O1xufTtcblxuZnVuY3Rpb24gY3NzV2l0aE1hcHBpbmdUb1N0cmluZyhpdGVtLCB1c2VTb3VyY2VNYXApIHtcbiAgdmFyIGNvbnRlbnQgPSBpdGVtWzFdIHx8ICcnOyAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgcHJlZmVyLWRlc3RydWN0dXJpbmdcblxuICB2YXIgY3NzTWFwcGluZyA9IGl0ZW1bM107XG5cbiAgaWYgKCFjc3NNYXBwaW5nKSB7XG4gICAgcmV0dXJuIGNvbnRlbnQ7XG4gIH1cblxuICBpZiAodXNlU291cmNlTWFwICYmIHR5cGVvZiBidG9hID09PSAnZnVuY3Rpb24nKSB7XG4gICAgdmFyIHNvdXJjZU1hcHBpbmcgPSB0b0NvbW1lbnQoY3NzTWFwcGluZyk7XG4gICAgdmFyIHNvdXJjZVVSTHMgPSBjc3NNYXBwaW5nLnNvdXJjZXMubWFwKGZ1bmN0aW9uIChzb3VyY2UpIHtcbiAgICAgIHJldHVybiBcIi8qIyBzb3VyY2VVUkw9XCIuY29uY2F0KGNzc01hcHBpbmcuc291cmNlUm9vdCB8fCAnJykuY29uY2F0KHNvdXJjZSwgXCIgKi9cIik7XG4gICAgfSk7XG4gICAgcmV0dXJuIFtjb250ZW50XS5jb25jYXQoc291cmNlVVJMcykuY29uY2F0KFtzb3VyY2VNYXBwaW5nXSkuam9pbignXFxuJyk7XG4gIH1cblxuICByZXR1cm4gW2NvbnRlbnRdLmpvaW4oJ1xcbicpO1xufSAvLyBBZGFwdGVkIGZyb20gY29udmVydC1zb3VyY2UtbWFwIChNSVQpXG5cblxuZnVuY3Rpb24gdG9Db21tZW50KHNvdXJjZU1hcCkge1xuICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgbm8tdW5kZWZcbiAgdmFyIGJhc2U2NCA9IGJ0b2EodW5lc2NhcGUoZW5jb2RlVVJJQ29tcG9uZW50KEpTT04uc3RyaW5naWZ5KHNvdXJjZU1hcCkpKSk7XG4gIHZhciBkYXRhID0gXCJzb3VyY2VNYXBwaW5nVVJMPWRhdGE6YXBwbGljYXRpb24vanNvbjtjaGFyc2V0PXV0Zi04O2Jhc2U2NCxcIi5jb25jYXQoYmFzZTY0KTtcbiAgcmV0dXJuIFwiLyojIFwiLmNvbmNhdChkYXRhLCBcIiAqL1wiKTtcbn1cblxuXG4vLy8vLy8vLy8vLy8vLy8vLy9cbi8vIFdFQlBBQ0sgRk9PVEVSXG4vLyAuL25vZGVfbW9kdWxlcy9jc3MtbG9hZGVyL2Rpc3QvcnVudGltZS9hcGkuanNcbi8vIG1vZHVsZSBpZCA9IDVcbi8vIG1vZHVsZSBjaHVua3MgPSAwIiwiLypcblx0TUlUIExpY2Vuc2UgaHR0cDovL3d3dy5vcGVuc291cmNlLm9yZy9saWNlbnNlcy9taXQtbGljZW5zZS5waHBcblx0QXV0aG9yIFRvYmlhcyBLb3BwZXJzIEBzb2tyYVxuKi9cblxudmFyIHN0eWxlc0luRG9tID0ge307XG5cbnZhclx0bWVtb2l6ZSA9IGZ1bmN0aW9uIChmbikge1xuXHR2YXIgbWVtbztcblxuXHRyZXR1cm4gZnVuY3Rpb24gKCkge1xuXHRcdGlmICh0eXBlb2YgbWVtbyA9PT0gXCJ1bmRlZmluZWRcIikgbWVtbyA9IGZuLmFwcGx5KHRoaXMsIGFyZ3VtZW50cyk7XG5cdFx0cmV0dXJuIG1lbW87XG5cdH07XG59O1xuXG52YXIgaXNPbGRJRSA9IG1lbW9pemUoZnVuY3Rpb24gKCkge1xuXHQvLyBUZXN0IGZvciBJRSA8PSA5IGFzIHByb3Bvc2VkIGJ5IEJyb3dzZXJoYWNrc1xuXHQvLyBAc2VlIGh0dHA6Ly9icm93c2VyaGFja3MuY29tLyNoYWNrLWU3MWQ4NjkyZjY1MzM0MTczZmVlNzE1YzIyMmNiODA1XG5cdC8vIFRlc3RzIGZvciBleGlzdGVuY2Ugb2Ygc3RhbmRhcmQgZ2xvYmFscyBpcyB0byBhbGxvdyBzdHlsZS1sb2FkZXJcblx0Ly8gdG8gb3BlcmF0ZSBjb3JyZWN0bHkgaW50byBub24tc3RhbmRhcmQgZW52aXJvbm1lbnRzXG5cdC8vIEBzZWUgaHR0cHM6Ly9naXRodWIuY29tL3dlYnBhY2stY29udHJpYi9zdHlsZS1sb2FkZXIvaXNzdWVzLzE3N1xuXHRyZXR1cm4gd2luZG93ICYmIGRvY3VtZW50ICYmIGRvY3VtZW50LmFsbCAmJiAhd2luZG93LmF0b2I7XG59KTtcblxudmFyIGdldEVsZW1lbnQgPSAoZnVuY3Rpb24gKGZuKSB7XG5cdHZhciBtZW1vID0ge307XG5cblx0cmV0dXJuIGZ1bmN0aW9uKHNlbGVjdG9yKSB7XG5cdFx0aWYgKHR5cGVvZiBtZW1vW3NlbGVjdG9yXSA9PT0gXCJ1bmRlZmluZWRcIikge1xuXHRcdFx0dmFyIHN0eWxlVGFyZ2V0ID0gZm4uY2FsbCh0aGlzLCBzZWxlY3Rvcik7XG5cdFx0XHQvLyBTcGVjaWFsIGNhc2UgdG8gcmV0dXJuIGhlYWQgb2YgaWZyYW1lIGluc3RlYWQgb2YgaWZyYW1lIGl0c2VsZlxuXHRcdFx0aWYgKHN0eWxlVGFyZ2V0IGluc3RhbmNlb2Ygd2luZG93LkhUTUxJRnJhbWVFbGVtZW50KSB7XG5cdFx0XHRcdHRyeSB7XG5cdFx0XHRcdFx0Ly8gVGhpcyB3aWxsIHRocm93IGFuIGV4Y2VwdGlvbiBpZiBhY2Nlc3MgdG8gaWZyYW1lIGlzIGJsb2NrZWRcblx0XHRcdFx0XHQvLyBkdWUgdG8gY3Jvc3Mtb3JpZ2luIHJlc3RyaWN0aW9uc1xuXHRcdFx0XHRcdHN0eWxlVGFyZ2V0ID0gc3R5bGVUYXJnZXQuY29udGVudERvY3VtZW50LmhlYWQ7XG5cdFx0XHRcdH0gY2F0Y2goZSkge1xuXHRcdFx0XHRcdHN0eWxlVGFyZ2V0ID0gbnVsbDtcblx0XHRcdFx0fVxuXHRcdFx0fVxuXHRcdFx0bWVtb1tzZWxlY3Rvcl0gPSBzdHlsZVRhcmdldDtcblx0XHR9XG5cdFx0cmV0dXJuIG1lbW9bc2VsZWN0b3JdXG5cdH07XG59KShmdW5jdGlvbiAodGFyZ2V0KSB7XG5cdHJldHVybiBkb2N1bWVudC5xdWVyeVNlbGVjdG9yKHRhcmdldClcbn0pO1xuXG52YXIgc2luZ2xldG9uID0gbnVsbDtcbnZhclx0c2luZ2xldG9uQ291bnRlciA9IDA7XG52YXJcdHN0eWxlc0luc2VydGVkQXRUb3AgPSBbXTtcblxudmFyXHRmaXhVcmxzID0gcmVxdWlyZShcIi4vdXJsc1wiKTtcblxubW9kdWxlLmV4cG9ydHMgPSBmdW5jdGlvbihsaXN0LCBvcHRpb25zKSB7XG5cdGlmICh0eXBlb2YgREVCVUcgIT09IFwidW5kZWZpbmVkXCIgJiYgREVCVUcpIHtcblx0XHRpZiAodHlwZW9mIGRvY3VtZW50ICE9PSBcIm9iamVjdFwiKSB0aHJvdyBuZXcgRXJyb3IoXCJUaGUgc3R5bGUtbG9hZGVyIGNhbm5vdCBiZSB1c2VkIGluIGEgbm9uLWJyb3dzZXIgZW52aXJvbm1lbnRcIik7XG5cdH1cblxuXHRvcHRpb25zID0gb3B0aW9ucyB8fCB7fTtcblxuXHRvcHRpb25zLmF0dHJzID0gdHlwZW9mIG9wdGlvbnMuYXR0cnMgPT09IFwib2JqZWN0XCIgPyBvcHRpb25zLmF0dHJzIDoge307XG5cblx0Ly8gRm9yY2Ugc2luZ2xlLXRhZyBzb2x1dGlvbiBvbiBJRTYtOSwgd2hpY2ggaGFzIGEgaGFyZCBsaW1pdCBvbiB0aGUgIyBvZiA8c3R5bGU+XG5cdC8vIHRhZ3MgaXQgd2lsbCBhbGxvdyBvbiBhIHBhZ2Vcblx0aWYgKCFvcHRpb25zLnNpbmdsZXRvbiAmJiB0eXBlb2Ygb3B0aW9ucy5zaW5nbGV0b24gIT09IFwiYm9vbGVhblwiKSBvcHRpb25zLnNpbmdsZXRvbiA9IGlzT2xkSUUoKTtcblxuXHQvLyBCeSBkZWZhdWx0LCBhZGQgPHN0eWxlPiB0YWdzIHRvIHRoZSA8aGVhZD4gZWxlbWVudFxuXHRpZiAoIW9wdGlvbnMuaW5zZXJ0SW50bykgb3B0aW9ucy5pbnNlcnRJbnRvID0gXCJoZWFkXCI7XG5cblx0Ly8gQnkgZGVmYXVsdCwgYWRkIDxzdHlsZT4gdGFncyB0byB0aGUgYm90dG9tIG9mIHRoZSB0YXJnZXRcblx0aWYgKCFvcHRpb25zLmluc2VydEF0KSBvcHRpb25zLmluc2VydEF0ID0gXCJib3R0b21cIjtcblxuXHR2YXIgc3R5bGVzID0gbGlzdFRvU3R5bGVzKGxpc3QsIG9wdGlvbnMpO1xuXG5cdGFkZFN0eWxlc1RvRG9tKHN0eWxlcywgb3B0aW9ucyk7XG5cblx0cmV0dXJuIGZ1bmN0aW9uIHVwZGF0ZSAobmV3TGlzdCkge1xuXHRcdHZhciBtYXlSZW1vdmUgPSBbXTtcblxuXHRcdGZvciAodmFyIGkgPSAwOyBpIDwgc3R5bGVzLmxlbmd0aDsgaSsrKSB7XG5cdFx0XHR2YXIgaXRlbSA9IHN0eWxlc1tpXTtcblx0XHRcdHZhciBkb21TdHlsZSA9IHN0eWxlc0luRG9tW2l0ZW0uaWRdO1xuXG5cdFx0XHRkb21TdHlsZS5yZWZzLS07XG5cdFx0XHRtYXlSZW1vdmUucHVzaChkb21TdHlsZSk7XG5cdFx0fVxuXG5cdFx0aWYobmV3TGlzdCkge1xuXHRcdFx0dmFyIG5ld1N0eWxlcyA9IGxpc3RUb1N0eWxlcyhuZXdMaXN0LCBvcHRpb25zKTtcblx0XHRcdGFkZFN0eWxlc1RvRG9tKG5ld1N0eWxlcywgb3B0aW9ucyk7XG5cdFx0fVxuXG5cdFx0Zm9yICh2YXIgaSA9IDA7IGkgPCBtYXlSZW1vdmUubGVuZ3RoOyBpKyspIHtcblx0XHRcdHZhciBkb21TdHlsZSA9IG1heVJlbW92ZVtpXTtcblxuXHRcdFx0aWYoZG9tU3R5bGUucmVmcyA9PT0gMCkge1xuXHRcdFx0XHRmb3IgKHZhciBqID0gMDsgaiA8IGRvbVN0eWxlLnBhcnRzLmxlbmd0aDsgaisrKSBkb21TdHlsZS5wYXJ0c1tqXSgpO1xuXG5cdFx0XHRcdGRlbGV0ZSBzdHlsZXNJbkRvbVtkb21TdHlsZS5pZF07XG5cdFx0XHR9XG5cdFx0fVxuXHR9O1xufTtcblxuZnVuY3Rpb24gYWRkU3R5bGVzVG9Eb20gKHN0eWxlcywgb3B0aW9ucykge1xuXHRmb3IgKHZhciBpID0gMDsgaSA8IHN0eWxlcy5sZW5ndGg7IGkrKykge1xuXHRcdHZhciBpdGVtID0gc3R5bGVzW2ldO1xuXHRcdHZhciBkb21TdHlsZSA9IHN0eWxlc0luRG9tW2l0ZW0uaWRdO1xuXG5cdFx0aWYoZG9tU3R5bGUpIHtcblx0XHRcdGRvbVN0eWxlLnJlZnMrKztcblxuXHRcdFx0Zm9yKHZhciBqID0gMDsgaiA8IGRvbVN0eWxlLnBhcnRzLmxlbmd0aDsgaisrKSB7XG5cdFx0XHRcdGRvbVN0eWxlLnBhcnRzW2pdKGl0ZW0ucGFydHNbal0pO1xuXHRcdFx0fVxuXG5cdFx0XHRmb3IoOyBqIDwgaXRlbS5wYXJ0cy5sZW5ndGg7IGorKykge1xuXHRcdFx0XHRkb21TdHlsZS5wYXJ0cy5wdXNoKGFkZFN0eWxlKGl0ZW0ucGFydHNbal0sIG9wdGlvbnMpKTtcblx0XHRcdH1cblx0XHR9IGVsc2Uge1xuXHRcdFx0dmFyIHBhcnRzID0gW107XG5cblx0XHRcdGZvcih2YXIgaiA9IDA7IGogPCBpdGVtLnBhcnRzLmxlbmd0aDsgaisrKSB7XG5cdFx0XHRcdHBhcnRzLnB1c2goYWRkU3R5bGUoaXRlbS5wYXJ0c1tqXSwgb3B0aW9ucykpO1xuXHRcdFx0fVxuXG5cdFx0XHRzdHlsZXNJbkRvbVtpdGVtLmlkXSA9IHtpZDogaXRlbS5pZCwgcmVmczogMSwgcGFydHM6IHBhcnRzfTtcblx0XHR9XG5cdH1cbn1cblxuZnVuY3Rpb24gbGlzdFRvU3R5bGVzIChsaXN0LCBvcHRpb25zKSB7XG5cdHZhciBzdHlsZXMgPSBbXTtcblx0dmFyIG5ld1N0eWxlcyA9IHt9O1xuXG5cdGZvciAodmFyIGkgPSAwOyBpIDwgbGlzdC5sZW5ndGg7IGkrKykge1xuXHRcdHZhciBpdGVtID0gbGlzdFtpXTtcblx0XHR2YXIgaWQgPSBvcHRpb25zLmJhc2UgPyBpdGVtWzBdICsgb3B0aW9ucy5iYXNlIDogaXRlbVswXTtcblx0XHR2YXIgY3NzID0gaXRlbVsxXTtcblx0XHR2YXIgbWVkaWEgPSBpdGVtWzJdO1xuXHRcdHZhciBzb3VyY2VNYXAgPSBpdGVtWzNdO1xuXHRcdHZhciBwYXJ0ID0ge2NzczogY3NzLCBtZWRpYTogbWVkaWEsIHNvdXJjZU1hcDogc291cmNlTWFwfTtcblxuXHRcdGlmKCFuZXdTdHlsZXNbaWRdKSBzdHlsZXMucHVzaChuZXdTdHlsZXNbaWRdID0ge2lkOiBpZCwgcGFydHM6IFtwYXJ0XX0pO1xuXHRcdGVsc2UgbmV3U3R5bGVzW2lkXS5wYXJ0cy5wdXNoKHBhcnQpO1xuXHR9XG5cblx0cmV0dXJuIHN0eWxlcztcbn1cblxuZnVuY3Rpb24gaW5zZXJ0U3R5bGVFbGVtZW50IChvcHRpb25zLCBzdHlsZSkge1xuXHR2YXIgdGFyZ2V0ID0gZ2V0RWxlbWVudChvcHRpb25zLmluc2VydEludG8pXG5cblx0aWYgKCF0YXJnZXQpIHtcblx0XHR0aHJvdyBuZXcgRXJyb3IoXCJDb3VsZG4ndCBmaW5kIGEgc3R5bGUgdGFyZ2V0LiBUaGlzIHByb2JhYmx5IG1lYW5zIHRoYXQgdGhlIHZhbHVlIGZvciB0aGUgJ2luc2VydEludG8nIHBhcmFtZXRlciBpcyBpbnZhbGlkLlwiKTtcblx0fVxuXG5cdHZhciBsYXN0U3R5bGVFbGVtZW50SW5zZXJ0ZWRBdFRvcCA9IHN0eWxlc0luc2VydGVkQXRUb3Bbc3R5bGVzSW5zZXJ0ZWRBdFRvcC5sZW5ndGggLSAxXTtcblxuXHRpZiAob3B0aW9ucy5pbnNlcnRBdCA9PT0gXCJ0b3BcIikge1xuXHRcdGlmICghbGFzdFN0eWxlRWxlbWVudEluc2VydGVkQXRUb3ApIHtcblx0XHRcdHRhcmdldC5pbnNlcnRCZWZvcmUoc3R5bGUsIHRhcmdldC5maXJzdENoaWxkKTtcblx0XHR9IGVsc2UgaWYgKGxhc3RTdHlsZUVsZW1lbnRJbnNlcnRlZEF0VG9wLm5leHRTaWJsaW5nKSB7XG5cdFx0XHR0YXJnZXQuaW5zZXJ0QmVmb3JlKHN0eWxlLCBsYXN0U3R5bGVFbGVtZW50SW5zZXJ0ZWRBdFRvcC5uZXh0U2libGluZyk7XG5cdFx0fSBlbHNlIHtcblx0XHRcdHRhcmdldC5hcHBlbmRDaGlsZChzdHlsZSk7XG5cdFx0fVxuXHRcdHN0eWxlc0luc2VydGVkQXRUb3AucHVzaChzdHlsZSk7XG5cdH0gZWxzZSBpZiAob3B0aW9ucy5pbnNlcnRBdCA9PT0gXCJib3R0b21cIikge1xuXHRcdHRhcmdldC5hcHBlbmRDaGlsZChzdHlsZSk7XG5cdH0gZWxzZSBpZiAodHlwZW9mIG9wdGlvbnMuaW5zZXJ0QXQgPT09IFwib2JqZWN0XCIgJiYgb3B0aW9ucy5pbnNlcnRBdC5iZWZvcmUpIHtcblx0XHR2YXIgbmV4dFNpYmxpbmcgPSBnZXRFbGVtZW50KG9wdGlvbnMuaW5zZXJ0SW50byArIFwiIFwiICsgb3B0aW9ucy5pbnNlcnRBdC5iZWZvcmUpO1xuXHRcdHRhcmdldC5pbnNlcnRCZWZvcmUoc3R5bGUsIG5leHRTaWJsaW5nKTtcblx0fSBlbHNlIHtcblx0XHR0aHJvdyBuZXcgRXJyb3IoXCJbU3R5bGUgTG9hZGVyXVxcblxcbiBJbnZhbGlkIHZhbHVlIGZvciBwYXJhbWV0ZXIgJ2luc2VydEF0JyAoJ29wdGlvbnMuaW5zZXJ0QXQnKSBmb3VuZC5cXG4gTXVzdCBiZSAndG9wJywgJ2JvdHRvbScsIG9yIE9iamVjdC5cXG4gKGh0dHBzOi8vZ2l0aHViLmNvbS93ZWJwYWNrLWNvbnRyaWIvc3R5bGUtbG9hZGVyI2luc2VydGF0KVxcblwiKTtcblx0fVxufVxuXG5mdW5jdGlvbiByZW1vdmVTdHlsZUVsZW1lbnQgKHN0eWxlKSB7XG5cdGlmIChzdHlsZS5wYXJlbnROb2RlID09PSBudWxsKSByZXR1cm4gZmFsc2U7XG5cdHN0eWxlLnBhcmVudE5vZGUucmVtb3ZlQ2hpbGQoc3R5bGUpO1xuXG5cdHZhciBpZHggPSBzdHlsZXNJbnNlcnRlZEF0VG9wLmluZGV4T2Yoc3R5bGUpO1xuXHRpZihpZHggPj0gMCkge1xuXHRcdHN0eWxlc0luc2VydGVkQXRUb3Auc3BsaWNlKGlkeCwgMSk7XG5cdH1cbn1cblxuZnVuY3Rpb24gY3JlYXRlU3R5bGVFbGVtZW50IChvcHRpb25zKSB7XG5cdHZhciBzdHlsZSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJzdHlsZVwiKTtcblxuXHRvcHRpb25zLmF0dHJzLnR5cGUgPSBcInRleHQvY3NzXCI7XG5cblx0YWRkQXR0cnMoc3R5bGUsIG9wdGlvbnMuYXR0cnMpO1xuXHRpbnNlcnRTdHlsZUVsZW1lbnQob3B0aW9ucywgc3R5bGUpO1xuXG5cdHJldHVybiBzdHlsZTtcbn1cblxuZnVuY3Rpb24gY3JlYXRlTGlua0VsZW1lbnQgKG9wdGlvbnMpIHtcblx0dmFyIGxpbmsgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwibGlua1wiKTtcblxuXHRvcHRpb25zLmF0dHJzLnR5cGUgPSBcInRleHQvY3NzXCI7XG5cdG9wdGlvbnMuYXR0cnMucmVsID0gXCJzdHlsZXNoZWV0XCI7XG5cblx0YWRkQXR0cnMobGluaywgb3B0aW9ucy5hdHRycyk7XG5cdGluc2VydFN0eWxlRWxlbWVudChvcHRpb25zLCBsaW5rKTtcblxuXHRyZXR1cm4gbGluaztcbn1cblxuZnVuY3Rpb24gYWRkQXR0cnMgKGVsLCBhdHRycykge1xuXHRPYmplY3Qua2V5cyhhdHRycykuZm9yRWFjaChmdW5jdGlvbiAoa2V5KSB7XG5cdFx0ZWwuc2V0QXR0cmlidXRlKGtleSwgYXR0cnNba2V5XSk7XG5cdH0pO1xufVxuXG5mdW5jdGlvbiBhZGRTdHlsZSAob2JqLCBvcHRpb25zKSB7XG5cdHZhciBzdHlsZSwgdXBkYXRlLCByZW1vdmUsIHJlc3VsdDtcblxuXHQvLyBJZiBhIHRyYW5zZm9ybSBmdW5jdGlvbiB3YXMgZGVmaW5lZCwgcnVuIGl0IG9uIHRoZSBjc3Ncblx0aWYgKG9wdGlvbnMudHJhbnNmb3JtICYmIG9iai5jc3MpIHtcblx0ICAgIHJlc3VsdCA9IG9wdGlvbnMudHJhbnNmb3JtKG9iai5jc3MpO1xuXG5cdCAgICBpZiAocmVzdWx0KSB7XG5cdCAgICBcdC8vIElmIHRyYW5zZm9ybSByZXR1cm5zIGEgdmFsdWUsIHVzZSB0aGF0IGluc3RlYWQgb2YgdGhlIG9yaWdpbmFsIGNzcy5cblx0ICAgIFx0Ly8gVGhpcyBhbGxvd3MgcnVubmluZyBydW50aW1lIHRyYW5zZm9ybWF0aW9ucyBvbiB0aGUgY3NzLlxuXHQgICAgXHRvYmouY3NzID0gcmVzdWx0O1xuXHQgICAgfSBlbHNlIHtcblx0ICAgIFx0Ly8gSWYgdGhlIHRyYW5zZm9ybSBmdW5jdGlvbiByZXR1cm5zIGEgZmFsc3kgdmFsdWUsIGRvbid0IGFkZCB0aGlzIGNzcy5cblx0ICAgIFx0Ly8gVGhpcyBhbGxvd3MgY29uZGl0aW9uYWwgbG9hZGluZyBvZiBjc3Ncblx0ICAgIFx0cmV0dXJuIGZ1bmN0aW9uKCkge1xuXHQgICAgXHRcdC8vIG5vb3Bcblx0ICAgIFx0fTtcblx0ICAgIH1cblx0fVxuXG5cdGlmIChvcHRpb25zLnNpbmdsZXRvbikge1xuXHRcdHZhciBzdHlsZUluZGV4ID0gc2luZ2xldG9uQ291bnRlcisrO1xuXG5cdFx0c3R5bGUgPSBzaW5nbGV0b24gfHwgKHNpbmdsZXRvbiA9IGNyZWF0ZVN0eWxlRWxlbWVudChvcHRpb25zKSk7XG5cblx0XHR1cGRhdGUgPSBhcHBseVRvU2luZ2xldG9uVGFnLmJpbmQobnVsbCwgc3R5bGUsIHN0eWxlSW5kZXgsIGZhbHNlKTtcblx0XHRyZW1vdmUgPSBhcHBseVRvU2luZ2xldG9uVGFnLmJpbmQobnVsbCwgc3R5bGUsIHN0eWxlSW5kZXgsIHRydWUpO1xuXG5cdH0gZWxzZSBpZiAoXG5cdFx0b2JqLnNvdXJjZU1hcCAmJlxuXHRcdHR5cGVvZiBVUkwgPT09IFwiZnVuY3Rpb25cIiAmJlxuXHRcdHR5cGVvZiBVUkwuY3JlYXRlT2JqZWN0VVJMID09PSBcImZ1bmN0aW9uXCIgJiZcblx0XHR0eXBlb2YgVVJMLnJldm9rZU9iamVjdFVSTCA9PT0gXCJmdW5jdGlvblwiICYmXG5cdFx0dHlwZW9mIEJsb2IgPT09IFwiZnVuY3Rpb25cIiAmJlxuXHRcdHR5cGVvZiBidG9hID09PSBcImZ1bmN0aW9uXCJcblx0KSB7XG5cdFx0c3R5bGUgPSBjcmVhdGVMaW5rRWxlbWVudChvcHRpb25zKTtcblx0XHR1cGRhdGUgPSB1cGRhdGVMaW5rLmJpbmQobnVsbCwgc3R5bGUsIG9wdGlvbnMpO1xuXHRcdHJlbW92ZSA9IGZ1bmN0aW9uICgpIHtcblx0XHRcdHJlbW92ZVN0eWxlRWxlbWVudChzdHlsZSk7XG5cblx0XHRcdGlmKHN0eWxlLmhyZWYpIFVSTC5yZXZva2VPYmplY3RVUkwoc3R5bGUuaHJlZik7XG5cdFx0fTtcblx0fSBlbHNlIHtcblx0XHRzdHlsZSA9IGNyZWF0ZVN0eWxlRWxlbWVudChvcHRpb25zKTtcblx0XHR1cGRhdGUgPSBhcHBseVRvVGFnLmJpbmQobnVsbCwgc3R5bGUpO1xuXHRcdHJlbW92ZSA9IGZ1bmN0aW9uICgpIHtcblx0XHRcdHJlbW92ZVN0eWxlRWxlbWVudChzdHlsZSk7XG5cdFx0fTtcblx0fVxuXG5cdHVwZGF0ZShvYmopO1xuXG5cdHJldHVybiBmdW5jdGlvbiB1cGRhdGVTdHlsZSAobmV3T2JqKSB7XG5cdFx0aWYgKG5ld09iaikge1xuXHRcdFx0aWYgKFxuXHRcdFx0XHRuZXdPYmouY3NzID09PSBvYmouY3NzICYmXG5cdFx0XHRcdG5ld09iai5tZWRpYSA9PT0gb2JqLm1lZGlhICYmXG5cdFx0XHRcdG5ld09iai5zb3VyY2VNYXAgPT09IG9iai5zb3VyY2VNYXBcblx0XHRcdCkge1xuXHRcdFx0XHRyZXR1cm47XG5cdFx0XHR9XG5cblx0XHRcdHVwZGF0ZShvYmogPSBuZXdPYmopO1xuXHRcdH0gZWxzZSB7XG5cdFx0XHRyZW1vdmUoKTtcblx0XHR9XG5cdH07XG59XG5cbnZhciByZXBsYWNlVGV4dCA9IChmdW5jdGlvbiAoKSB7XG5cdHZhciB0ZXh0U3RvcmUgPSBbXTtcblxuXHRyZXR1cm4gZnVuY3Rpb24gKGluZGV4LCByZXBsYWNlbWVudCkge1xuXHRcdHRleHRTdG9yZVtpbmRleF0gPSByZXBsYWNlbWVudDtcblxuXHRcdHJldHVybiB0ZXh0U3RvcmUuZmlsdGVyKEJvb2xlYW4pLmpvaW4oJ1xcbicpO1xuXHR9O1xufSkoKTtcblxuZnVuY3Rpb24gYXBwbHlUb1NpbmdsZXRvblRhZyAoc3R5bGUsIGluZGV4LCByZW1vdmUsIG9iaikge1xuXHR2YXIgY3NzID0gcmVtb3ZlID8gXCJcIiA6IG9iai5jc3M7XG5cblx0aWYgKHN0eWxlLnN0eWxlU2hlZXQpIHtcblx0XHRzdHlsZS5zdHlsZVNoZWV0LmNzc1RleHQgPSByZXBsYWNlVGV4dChpbmRleCwgY3NzKTtcblx0fSBlbHNlIHtcblx0XHR2YXIgY3NzTm9kZSA9IGRvY3VtZW50LmNyZWF0ZVRleHROb2RlKGNzcyk7XG5cdFx0dmFyIGNoaWxkTm9kZXMgPSBzdHlsZS5jaGlsZE5vZGVzO1xuXG5cdFx0aWYgKGNoaWxkTm9kZXNbaW5kZXhdKSBzdHlsZS5yZW1vdmVDaGlsZChjaGlsZE5vZGVzW2luZGV4XSk7XG5cblx0XHRpZiAoY2hpbGROb2Rlcy5sZW5ndGgpIHtcblx0XHRcdHN0eWxlLmluc2VydEJlZm9yZShjc3NOb2RlLCBjaGlsZE5vZGVzW2luZGV4XSk7XG5cdFx0fSBlbHNlIHtcblx0XHRcdHN0eWxlLmFwcGVuZENoaWxkKGNzc05vZGUpO1xuXHRcdH1cblx0fVxufVxuXG5mdW5jdGlvbiBhcHBseVRvVGFnIChzdHlsZSwgb2JqKSB7XG5cdHZhciBjc3MgPSBvYmouY3NzO1xuXHR2YXIgbWVkaWEgPSBvYmoubWVkaWE7XG5cblx0aWYobWVkaWEpIHtcblx0XHRzdHlsZS5zZXRBdHRyaWJ1dGUoXCJtZWRpYVwiLCBtZWRpYSlcblx0fVxuXG5cdGlmKHN0eWxlLnN0eWxlU2hlZXQpIHtcblx0XHRzdHlsZS5zdHlsZVNoZWV0LmNzc1RleHQgPSBjc3M7XG5cdH0gZWxzZSB7XG5cdFx0d2hpbGUoc3R5bGUuZmlyc3RDaGlsZCkge1xuXHRcdFx0c3R5bGUucmVtb3ZlQ2hpbGQoc3R5bGUuZmlyc3RDaGlsZCk7XG5cdFx0fVxuXG5cdFx0c3R5bGUuYXBwZW5kQ2hpbGQoZG9jdW1lbnQuY3JlYXRlVGV4dE5vZGUoY3NzKSk7XG5cdH1cbn1cblxuZnVuY3Rpb24gdXBkYXRlTGluayAobGluaywgb3B0aW9ucywgb2JqKSB7XG5cdHZhciBjc3MgPSBvYmouY3NzO1xuXHR2YXIgc291cmNlTWFwID0gb2JqLnNvdXJjZU1hcDtcblxuXHQvKlxuXHRcdElmIGNvbnZlcnRUb0Fic29sdXRlVXJscyBpc24ndCBkZWZpbmVkLCBidXQgc291cmNlbWFwcyBhcmUgZW5hYmxlZFxuXHRcdGFuZCB0aGVyZSBpcyBubyBwdWJsaWNQYXRoIGRlZmluZWQgdGhlbiBsZXRzIHR1cm4gY29udmVydFRvQWJzb2x1dGVVcmxzXG5cdFx0b24gYnkgZGVmYXVsdC4gIE90aGVyd2lzZSBkZWZhdWx0IHRvIHRoZSBjb252ZXJ0VG9BYnNvbHV0ZVVybHMgb3B0aW9uXG5cdFx0ZGlyZWN0bHlcblx0Ki9cblx0dmFyIGF1dG9GaXhVcmxzID0gb3B0aW9ucy5jb252ZXJ0VG9BYnNvbHV0ZVVybHMgPT09IHVuZGVmaW5lZCAmJiBzb3VyY2VNYXA7XG5cblx0aWYgKG9wdGlvbnMuY29udmVydFRvQWJzb2x1dGVVcmxzIHx8IGF1dG9GaXhVcmxzKSB7XG5cdFx0Y3NzID0gZml4VXJscyhjc3MpO1xuXHR9XG5cblx0aWYgKHNvdXJjZU1hcCkge1xuXHRcdC8vIGh0dHA6Ly9zdGFja292ZXJmbG93LmNvbS9hLzI2NjAzODc1XG5cdFx0Y3NzICs9IFwiXFxuLyojIHNvdXJjZU1hcHBpbmdVUkw9ZGF0YTphcHBsaWNhdGlvbi9qc29uO2Jhc2U2NCxcIiArIGJ0b2EodW5lc2NhcGUoZW5jb2RlVVJJQ29tcG9uZW50KEpTT04uc3RyaW5naWZ5KHNvdXJjZU1hcCkpKSkgKyBcIiAqL1wiO1xuXHR9XG5cblx0dmFyIGJsb2IgPSBuZXcgQmxvYihbY3NzXSwgeyB0eXBlOiBcInRleHQvY3NzXCIgfSk7XG5cblx0dmFyIG9sZFNyYyA9IGxpbmsuaHJlZjtcblxuXHRsaW5rLmhyZWYgPSBVUkwuY3JlYXRlT2JqZWN0VVJMKGJsb2IpO1xuXG5cdGlmKG9sZFNyYykgVVJMLnJldm9rZU9iamVjdFVSTChvbGRTcmMpO1xufVxuXG5cblxuLy8vLy8vLy8vLy8vLy8vLy8vXG4vLyBXRUJQQUNLIEZPT1RFUlxuLy8gLi9ub2RlX21vZHVsZXMvc3R5bGUtbG9hZGVyL2xpYi9hZGRTdHlsZXMuanNcbi8vIG1vZHVsZSBpZCA9IDZcbi8vIG1vZHVsZSBjaHVua3MgPSAwIiwiXG4vKipcbiAqIFdoZW4gc291cmNlIG1hcHMgYXJlIGVuYWJsZWQsIGBzdHlsZS1sb2FkZXJgIHVzZXMgYSBsaW5rIGVsZW1lbnQgd2l0aCBhIGRhdGEtdXJpIHRvXG4gKiBlbWJlZCB0aGUgY3NzIG9uIHRoZSBwYWdlLiBUaGlzIGJyZWFrcyBhbGwgcmVsYXRpdmUgdXJscyBiZWNhdXNlIG5vdyB0aGV5IGFyZSByZWxhdGl2ZSB0byBhXG4gKiBidW5kbGUgaW5zdGVhZCBvZiB0aGUgY3VycmVudCBwYWdlLlxuICpcbiAqIE9uZSBzb2x1dGlvbiBpcyB0byBvbmx5IHVzZSBmdWxsIHVybHMsIGJ1dCB0aGF0IG1heSBiZSBpbXBvc3NpYmxlLlxuICpcbiAqIEluc3RlYWQsIHRoaXMgZnVuY3Rpb24gXCJmaXhlc1wiIHRoZSByZWxhdGl2ZSB1cmxzIHRvIGJlIGFic29sdXRlIGFjY29yZGluZyB0byB0aGUgY3VycmVudCBwYWdlIGxvY2F0aW9uLlxuICpcbiAqIEEgcnVkaW1lbnRhcnkgdGVzdCBzdWl0ZSBpcyBsb2NhdGVkIGF0IGB0ZXN0L2ZpeFVybHMuanNgIGFuZCBjYW4gYmUgcnVuIHZpYSB0aGUgYG5wbSB0ZXN0YCBjb21tYW5kLlxuICpcbiAqL1xuXG5tb2R1bGUuZXhwb3J0cyA9IGZ1bmN0aW9uIChjc3MpIHtcbiAgLy8gZ2V0IGN1cnJlbnQgbG9jYXRpb25cbiAgdmFyIGxvY2F0aW9uID0gdHlwZW9mIHdpbmRvdyAhPT0gXCJ1bmRlZmluZWRcIiAmJiB3aW5kb3cubG9jYXRpb247XG5cbiAgaWYgKCFsb2NhdGlvbikge1xuICAgIHRocm93IG5ldyBFcnJvcihcImZpeFVybHMgcmVxdWlyZXMgd2luZG93LmxvY2F0aW9uXCIpO1xuICB9XG5cblx0Ly8gYmxhbmsgb3IgbnVsbD9cblx0aWYgKCFjc3MgfHwgdHlwZW9mIGNzcyAhPT0gXCJzdHJpbmdcIikge1xuXHQgIHJldHVybiBjc3M7XG4gIH1cblxuICB2YXIgYmFzZVVybCA9IGxvY2F0aW9uLnByb3RvY29sICsgXCIvL1wiICsgbG9jYXRpb24uaG9zdDtcbiAgdmFyIGN1cnJlbnREaXIgPSBiYXNlVXJsICsgbG9jYXRpb24ucGF0aG5hbWUucmVwbGFjZSgvXFwvW15cXC9dKiQvLCBcIi9cIik7XG5cblx0Ly8gY29udmVydCBlYWNoIHVybCguLi4pXG5cdC8qXG5cdFRoaXMgcmVndWxhciBleHByZXNzaW9uIGlzIGp1c3QgYSB3YXkgdG8gcmVjdXJzaXZlbHkgbWF0Y2ggYnJhY2tldHMgd2l0aGluXG5cdGEgc3RyaW5nLlxuXG5cdCAvdXJsXFxzKlxcKCAgPSBNYXRjaCBvbiB0aGUgd29yZCBcInVybFwiIHdpdGggYW55IHdoaXRlc3BhY2UgYWZ0ZXIgaXQgYW5kIHRoZW4gYSBwYXJlbnNcblx0ICAgKCAgPSBTdGFydCBhIGNhcHR1cmluZyBncm91cFxuXHQgICAgICg/OiAgPSBTdGFydCBhIG5vbi1jYXB0dXJpbmcgZ3JvdXBcblx0ICAgICAgICAgW14pKF0gID0gTWF0Y2ggYW55dGhpbmcgdGhhdCBpc24ndCBhIHBhcmVudGhlc2VzXG5cdCAgICAgICAgIHwgID0gT1Jcblx0ICAgICAgICAgXFwoICA9IE1hdGNoIGEgc3RhcnQgcGFyZW50aGVzZXNcblx0ICAgICAgICAgICAgICg/OiAgPSBTdGFydCBhbm90aGVyIG5vbi1jYXB0dXJpbmcgZ3JvdXBzXG5cdCAgICAgICAgICAgICAgICAgW14pKF0rICA9IE1hdGNoIGFueXRoaW5nIHRoYXQgaXNuJ3QgYSBwYXJlbnRoZXNlc1xuXHQgICAgICAgICAgICAgICAgIHwgID0gT1Jcblx0ICAgICAgICAgICAgICAgICBcXCggID0gTWF0Y2ggYSBzdGFydCBwYXJlbnRoZXNlc1xuXHQgICAgICAgICAgICAgICAgICAgICBbXikoXSogID0gTWF0Y2ggYW55dGhpbmcgdGhhdCBpc24ndCBhIHBhcmVudGhlc2VzXG5cdCAgICAgICAgICAgICAgICAgXFwpICA9IE1hdGNoIGEgZW5kIHBhcmVudGhlc2VzXG5cdCAgICAgICAgICAgICApICA9IEVuZCBHcm91cFxuICAgICAgICAgICAgICAqXFwpID0gTWF0Y2ggYW55dGhpbmcgYW5kIHRoZW4gYSBjbG9zZSBwYXJlbnNcbiAgICAgICAgICApICA9IENsb3NlIG5vbi1jYXB0dXJpbmcgZ3JvdXBcbiAgICAgICAgICAqICA9IE1hdGNoIGFueXRoaW5nXG4gICAgICAgKSAgPSBDbG9zZSBjYXB0dXJpbmcgZ3JvdXBcblx0IFxcKSAgPSBNYXRjaCBhIGNsb3NlIHBhcmVuc1xuXG5cdCAvZ2kgID0gR2V0IGFsbCBtYXRjaGVzLCBub3QgdGhlIGZpcnN0LiAgQmUgY2FzZSBpbnNlbnNpdGl2ZS5cblx0ICovXG5cdHZhciBmaXhlZENzcyA9IGNzcy5yZXBsYWNlKC91cmxcXHMqXFwoKCg/OlteKShdfFxcKCg/OlteKShdK3xcXChbXikoXSpcXCkpKlxcKSkqKVxcKS9naSwgZnVuY3Rpb24oZnVsbE1hdGNoLCBvcmlnVXJsKSB7XG5cdFx0Ly8gc3RyaXAgcXVvdGVzIChpZiB0aGV5IGV4aXN0KVxuXHRcdHZhciB1bnF1b3RlZE9yaWdVcmwgPSBvcmlnVXJsXG5cdFx0XHQudHJpbSgpXG5cdFx0XHQucmVwbGFjZSgvXlwiKC4qKVwiJC8sIGZ1bmN0aW9uKG8sICQxKXsgcmV0dXJuICQxOyB9KVxuXHRcdFx0LnJlcGxhY2UoL14nKC4qKSckLywgZnVuY3Rpb24obywgJDEpeyByZXR1cm4gJDE7IH0pO1xuXG5cdFx0Ly8gYWxyZWFkeSBhIGZ1bGwgdXJsPyBubyBjaGFuZ2Vcblx0XHRpZiAoL14oI3xkYXRhOnxodHRwOlxcL1xcL3xodHRwczpcXC9cXC98ZmlsZTpcXC9cXC9cXC8pL2kudGVzdCh1bnF1b3RlZE9yaWdVcmwpKSB7XG5cdFx0ICByZXR1cm4gZnVsbE1hdGNoO1xuXHRcdH1cblxuXHRcdC8vIGNvbnZlcnQgdGhlIHVybCB0byBhIGZ1bGwgdXJsXG5cdFx0dmFyIG5ld1VybDtcblxuXHRcdGlmICh1bnF1b3RlZE9yaWdVcmwuaW5kZXhPZihcIi8vXCIpID09PSAwKSB7XG5cdFx0ICBcdC8vVE9ETzogc2hvdWxkIHdlIGFkZCBwcm90b2NvbD9cblx0XHRcdG5ld1VybCA9IHVucXVvdGVkT3JpZ1VybDtcblx0XHR9IGVsc2UgaWYgKHVucXVvdGVkT3JpZ1VybC5pbmRleE9mKFwiL1wiKSA9PT0gMCkge1xuXHRcdFx0Ly8gcGF0aCBzaG91bGQgYmUgcmVsYXRpdmUgdG8gdGhlIGJhc2UgdXJsXG5cdFx0XHRuZXdVcmwgPSBiYXNlVXJsICsgdW5xdW90ZWRPcmlnVXJsOyAvLyBhbHJlYWR5IHN0YXJ0cyB3aXRoICcvJ1xuXHRcdH0gZWxzZSB7XG5cdFx0XHQvLyBwYXRoIHNob3VsZCBiZSByZWxhdGl2ZSB0byBjdXJyZW50IGRpcmVjdG9yeVxuXHRcdFx0bmV3VXJsID0gY3VycmVudERpciArIHVucXVvdGVkT3JpZ1VybC5yZXBsYWNlKC9eXFwuXFwvLywgXCJcIik7IC8vIFN0cmlwIGxlYWRpbmcgJy4vJ1xuXHRcdH1cblxuXHRcdC8vIHNlbmQgYmFjayB0aGUgZml4ZWQgdXJsKC4uLilcblx0XHRyZXR1cm4gXCJ1cmwoXCIgKyBKU09OLnN0cmluZ2lmeShuZXdVcmwpICsgXCIpXCI7XG5cdH0pO1xuXG5cdC8vIHNlbmQgYmFjayB0aGUgZml4ZWQgY3NzXG5cdHJldHVybiBmaXhlZENzcztcbn07XG5cblxuXG4vLy8vLy8vLy8vLy8vLy8vLy9cbi8vIFdFQlBBQ0sgRk9PVEVSXG4vLyAuL25vZGVfbW9kdWxlcy9zdHlsZS1sb2FkZXIvbGliL3VybHMuanNcbi8vIG1vZHVsZSBpZCA9IDdcbi8vIG1vZHVsZSBjaHVua3MgPSAwIiwiLy8gQ29weXJpZ2h0IDIwMTggVGhlIEdvIEF1dGhvcnMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuXG4vLyBVc2Ugb2YgdGhpcyBzb3VyY2UgY29kZSBpcyBnb3Zlcm5lZCBieSBhIEJTRC1zdHlsZVxuLy8gbGljZW5zZSB0aGF0IGNhbiBiZSBmb3VuZCBpbiB0aGUgTElDRU5TRSBmaWxlLlxuXG5cInVzZSBzdHJpY3RcIjtcblxuKCgpID0+IHtcblx0Y29uc3QgZW5vc3lzID0gKCkgPT4ge1xuXHRcdGNvbnN0IGVyciA9IG5ldyBFcnJvcihcIm5vdCBpbXBsZW1lbnRlZFwiKTtcblx0XHRlcnIuY29kZSA9IFwiRU5PU1lTXCI7XG5cdFx0cmV0dXJuIGVycjtcblx0fTtcblxuXHRpZiAoIWdsb2JhbFRoaXMuZnMpIHtcblx0XHRsZXQgb3V0cHV0QnVmID0gXCJcIjtcblx0XHRnbG9iYWxUaGlzLmZzID0ge1xuXHRcdFx0Y29uc3RhbnRzOiB7IE9fV1JPTkxZOiAtMSwgT19SRFdSOiAtMSwgT19DUkVBVDogLTEsIE9fVFJVTkM6IC0xLCBPX0FQUEVORDogLTEsIE9fRVhDTDogLTEgfSwgLy8gdW51c2VkXG5cdFx0XHR3cml0ZVN5bmMoZmQsIGJ1Zikge1xuXHRcdFx0XHRvdXRwdXRCdWYgKz0gZGVjb2Rlci5kZWNvZGUoYnVmKTtcblx0XHRcdFx0Y29uc3QgbmwgPSBvdXRwdXRCdWYubGFzdEluZGV4T2YoXCJcXG5cIik7XG5cdFx0XHRcdGlmIChubCAhPSAtMSkge1xuXHRcdFx0XHRcdGNvbnNvbGUubG9nKG91dHB1dEJ1Zi5zdWJzdHJpbmcoMCwgbmwpKTtcblx0XHRcdFx0XHRvdXRwdXRCdWYgPSBvdXRwdXRCdWYuc3Vic3RyaW5nKG5sICsgMSk7XG5cdFx0XHRcdH1cblx0XHRcdFx0cmV0dXJuIGJ1Zi5sZW5ndGg7XG5cdFx0XHR9LFxuXHRcdFx0d3JpdGUoZmQsIGJ1Ziwgb2Zmc2V0LCBsZW5ndGgsIHBvc2l0aW9uLCBjYWxsYmFjaykge1xuXHRcdFx0XHRpZiAob2Zmc2V0ICE9PSAwIHx8IGxlbmd0aCAhPT0gYnVmLmxlbmd0aCB8fCBwb3NpdGlvbiAhPT0gbnVsbCkge1xuXHRcdFx0XHRcdGNhbGxiYWNrKGVub3N5cygpKTtcblx0XHRcdFx0XHRyZXR1cm47XG5cdFx0XHRcdH1cblx0XHRcdFx0Y29uc3QgbiA9IHRoaXMud3JpdGVTeW5jKGZkLCBidWYpO1xuXHRcdFx0XHRjYWxsYmFjayhudWxsLCBuKTtcblx0XHRcdH0sXG5cdFx0XHRjaG1vZChwYXRoLCBtb2RlLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRjaG93bihwYXRoLCB1aWQsIGdpZCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0Y2xvc2UoZmQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGZjaG1vZChmZCwgbW9kZSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0ZmNob3duKGZkLCB1aWQsIGdpZCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0ZnN0YXQoZmQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGZzeW5jKGZkLCBjYWxsYmFjaykgeyBjYWxsYmFjayhudWxsKTsgfSxcblx0XHRcdGZ0cnVuY2F0ZShmZCwgbGVuZ3RoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRsY2hvd24ocGF0aCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdGxpbmsocGF0aCwgbGluaywgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0bHN0YXQocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0bWtkaXIocGF0aCwgcGVybSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0b3BlbihwYXRoLCBmbGFncywgbW9kZSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0cmVhZChmZCwgYnVmZmVyLCBvZmZzZXQsIGxlbmd0aCwgcG9zaXRpb24sIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHJlYWRkaXIocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0cmVhZGxpbmsocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0cmVuYW1lKGZyb20sIHRvLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRybWRpcihwYXRoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHRzdGF0KHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSxcblx0XHRcdHN5bWxpbmsocGF0aCwgbGluaywgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0dHJ1bmNhdGUocGF0aCwgbGVuZ3RoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sXG5cdFx0XHR1bmxpbmsocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdFx0dXRpbWVzKHBhdGgsIGF0aW1lLCBtdGltZSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LFxuXHRcdH07XG5cdH1cblxuXHRpZiAoIWdsb2JhbFRoaXMucHJvY2Vzcykge1xuXHRcdGdsb2JhbFRoaXMucHJvY2VzcyA9IHtcblx0XHRcdGdldHVpZCgpIHsgcmV0dXJuIC0xOyB9LFxuXHRcdFx0Z2V0Z2lkKCkgeyByZXR1cm4gLTE7IH0sXG5cdFx0XHRnZXRldWlkKCkgeyByZXR1cm4gLTE7IH0sXG5cdFx0XHRnZXRlZ2lkKCkgeyByZXR1cm4gLTE7IH0sXG5cdFx0XHRnZXRncm91cHMoKSB7IHRocm93IGVub3N5cygpOyB9LFxuXHRcdFx0cGlkOiAtMSxcblx0XHRcdHBwaWQ6IC0xLFxuXHRcdFx0dW1hc2soKSB7IHRocm93IGVub3N5cygpOyB9LFxuXHRcdFx0Y3dkKCkgeyB0aHJvdyBlbm9zeXMoKTsgfSxcblx0XHRcdGNoZGlyKCkgeyB0aHJvdyBlbm9zeXMoKTsgfSxcblx0XHR9XG5cdH1cblxuXHRpZiAoIWdsb2JhbFRoaXMuY3J5cHRvKSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiZ2xvYmFsVGhpcy5jcnlwdG8gaXMgbm90IGF2YWlsYWJsZSwgcG9seWZpbGwgcmVxdWlyZWQgKGNyeXB0by5nZXRSYW5kb21WYWx1ZXMgb25seSlcIik7XG5cdH1cblxuXHRpZiAoIWdsb2JhbFRoaXMucGVyZm9ybWFuY2UpIHtcblx0XHR0aHJvdyBuZXcgRXJyb3IoXCJnbG9iYWxUaGlzLnBlcmZvcm1hbmNlIGlzIG5vdCBhdmFpbGFibGUsIHBvbHlmaWxsIHJlcXVpcmVkIChwZXJmb3JtYW5jZS5ub3cgb25seSlcIik7XG5cdH1cblxuXHRpZiAoIWdsb2JhbFRoaXMuVGV4dEVuY29kZXIpIHtcblx0XHR0aHJvdyBuZXcgRXJyb3IoXCJnbG9iYWxUaGlzLlRleHRFbmNvZGVyIGlzIG5vdCBhdmFpbGFibGUsIHBvbHlmaWxsIHJlcXVpcmVkXCIpO1xuXHR9XG5cblx0aWYgKCFnbG9iYWxUaGlzLlRleHREZWNvZGVyKSB7XG5cdFx0dGhyb3cgbmV3IEVycm9yKFwiZ2xvYmFsVGhpcy5UZXh0RGVjb2RlciBpcyBub3QgYXZhaWxhYmxlLCBwb2x5ZmlsbCByZXF1aXJlZFwiKTtcblx0fVxuXG5cdGNvbnN0IGVuY29kZXIgPSBuZXcgVGV4dEVuY29kZXIoXCJ1dGYtOFwiKTtcblx0Y29uc3QgZGVjb2RlciA9IG5ldyBUZXh0RGVjb2RlcihcInV0Zi04XCIpO1xuXG5cdGdsb2JhbFRoaXMuR28gPSBjbGFzcyB7XG5cdFx0Y29uc3RydWN0b3IoKSB7XG5cdFx0XHR0aGlzLmFyZ3YgPSBbXCJqc1wiXTtcblx0XHRcdHRoaXMuZW52ID0ge307XG5cdFx0XHR0aGlzLmV4aXQgPSAoY29kZSkgPT4ge1xuXHRcdFx0XHRpZiAoY29kZSAhPT0gMCkge1xuXHRcdFx0XHRcdGNvbnNvbGUud2FybihcImV4aXQgY29kZTpcIiwgY29kZSk7XG5cdFx0XHRcdH1cblx0XHRcdH07XG5cdFx0XHR0aGlzLl9leGl0UHJvbWlzZSA9IG5ldyBQcm9taXNlKChyZXNvbHZlKSA9PiB7XG5cdFx0XHRcdHRoaXMuX3Jlc29sdmVFeGl0UHJvbWlzZSA9IHJlc29sdmU7XG5cdFx0XHR9KTtcblx0XHRcdHRoaXMuX3BlbmRpbmdFdmVudCA9IG51bGw7XG5cdFx0XHR0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cyA9IG5ldyBNYXAoKTtcblx0XHRcdHRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRCA9IDE7XG5cblx0XHRcdGNvbnN0IHNldEludDY0ID0gKGFkZHIsIHYpID0+IHtcblx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDMyKGFkZHIgKyAwLCB2LCB0cnVlKTtcblx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDMyKGFkZHIgKyA0LCBNYXRoLmZsb29yKHYgLyA0Mjk0OTY3Mjk2KSwgdHJ1ZSk7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IHNldEludDMyID0gKGFkZHIsIHYpID0+IHtcblx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDMyKGFkZHIgKyAwLCB2LCB0cnVlKTtcblx0XHRcdH1cblxuXHRcdFx0Y29uc3QgZ2V0SW50NjQgPSAoYWRkcikgPT4ge1xuXHRcdFx0XHRjb25zdCBsb3cgPSB0aGlzLm1lbS5nZXRVaW50MzIoYWRkciArIDAsIHRydWUpO1xuXHRcdFx0XHRjb25zdCBoaWdoID0gdGhpcy5tZW0uZ2V0SW50MzIoYWRkciArIDQsIHRydWUpO1xuXHRcdFx0XHRyZXR1cm4gbG93ICsgaGlnaCAqIDQyOTQ5NjcyOTY7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IGxvYWRWYWx1ZSA9IChhZGRyKSA9PiB7XG5cdFx0XHRcdGNvbnN0IGYgPSB0aGlzLm1lbS5nZXRGbG9hdDY0KGFkZHIsIHRydWUpO1xuXHRcdFx0XHRpZiAoZiA9PT0gMCkge1xuXHRcdFx0XHRcdHJldHVybiB1bmRlZmluZWQ7XG5cdFx0XHRcdH1cblx0XHRcdFx0aWYgKCFpc05hTihmKSkge1xuXHRcdFx0XHRcdHJldHVybiBmO1xuXHRcdFx0XHR9XG5cblx0XHRcdFx0Y29uc3QgaWQgPSB0aGlzLm1lbS5nZXRVaW50MzIoYWRkciwgdHJ1ZSk7XG5cdFx0XHRcdHJldHVybiB0aGlzLl92YWx1ZXNbaWRdO1xuXHRcdFx0fVxuXG5cdFx0XHRjb25zdCBzdG9yZVZhbHVlID0gKGFkZHIsIHYpID0+IHtcblx0XHRcdFx0Y29uc3QgbmFuSGVhZCA9IDB4N0ZGODAwMDA7XG5cblx0XHRcdFx0aWYgKHR5cGVvZiB2ID09PSBcIm51bWJlclwiICYmIHYgIT09IDApIHtcblx0XHRcdFx0XHRpZiAoaXNOYU4odikpIHtcblx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihhZGRyICsgNCwgbmFuSGVhZCwgdHJ1ZSk7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIoYWRkciwgMCwgdHJ1ZSk7XG5cdFx0XHRcdFx0XHRyZXR1cm47XG5cdFx0XHRcdFx0fVxuXHRcdFx0XHRcdHRoaXMubWVtLnNldEZsb2F0NjQoYWRkciwgdiwgdHJ1ZSk7XG5cdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHR9XG5cblx0XHRcdFx0aWYgKHYgPT09IHVuZGVmaW5lZCkge1xuXHRcdFx0XHRcdHRoaXMubWVtLnNldEZsb2F0NjQoYWRkciwgMCwgdHJ1ZSk7XG5cdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHR9XG5cblx0XHRcdFx0bGV0IGlkID0gdGhpcy5faWRzLmdldCh2KTtcblx0XHRcdFx0aWYgKGlkID09PSB1bmRlZmluZWQpIHtcblx0XHRcdFx0XHRpZCA9IHRoaXMuX2lkUG9vbC5wb3AoKTtcblx0XHRcdFx0XHRpZiAoaWQgPT09IHVuZGVmaW5lZCkge1xuXHRcdFx0XHRcdFx0aWQgPSB0aGlzLl92YWx1ZXMubGVuZ3RoO1xuXHRcdFx0XHRcdH1cblx0XHRcdFx0XHR0aGlzLl92YWx1ZXNbaWRdID0gdjtcblx0XHRcdFx0XHR0aGlzLl9nb1JlZkNvdW50c1tpZF0gPSAwO1xuXHRcdFx0XHRcdHRoaXMuX2lkcy5zZXQodiwgaWQpO1xuXHRcdFx0XHR9XG5cdFx0XHRcdHRoaXMuX2dvUmVmQ291bnRzW2lkXSsrO1xuXHRcdFx0XHRsZXQgdHlwZUZsYWcgPSAwO1xuXHRcdFx0XHRzd2l0Y2ggKHR5cGVvZiB2KSB7XG5cdFx0XHRcdFx0Y2FzZSBcIm9iamVjdFwiOlxuXHRcdFx0XHRcdFx0aWYgKHYgIT09IG51bGwpIHtcblx0XHRcdFx0XHRcdFx0dHlwZUZsYWcgPSAxO1xuXHRcdFx0XHRcdFx0fVxuXHRcdFx0XHRcdFx0YnJlYWs7XG5cdFx0XHRcdFx0Y2FzZSBcInN0cmluZ1wiOlxuXHRcdFx0XHRcdFx0dHlwZUZsYWcgPSAyO1xuXHRcdFx0XHRcdFx0YnJlYWs7XG5cdFx0XHRcdFx0Y2FzZSBcInN5bWJvbFwiOlxuXHRcdFx0XHRcdFx0dHlwZUZsYWcgPSAzO1xuXHRcdFx0XHRcdFx0YnJlYWs7XG5cdFx0XHRcdFx0Y2FzZSBcImZ1bmN0aW9uXCI6XG5cdFx0XHRcdFx0XHR0eXBlRmxhZyA9IDQ7XG5cdFx0XHRcdFx0XHRicmVhaztcblx0XHRcdFx0fVxuXHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQgfCB0eXBlRmxhZywgdHJ1ZSk7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihhZGRyLCBpZCwgdHJ1ZSk7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IGxvYWRTbGljZSA9IChhZGRyKSA9PiB7XG5cdFx0XHRcdGNvbnN0IGFycmF5ID0gZ2V0SW50NjQoYWRkciArIDApO1xuXHRcdFx0XHRjb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7XG5cdFx0XHRcdHJldHVybiBuZXcgVWludDhBcnJheSh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlciwgYXJyYXksIGxlbik7XG5cdFx0XHR9XG5cblx0XHRcdGNvbnN0IGxvYWRTbGljZU9mVmFsdWVzID0gKGFkZHIpID0+IHtcblx0XHRcdFx0Y29uc3QgYXJyYXkgPSBnZXRJbnQ2NChhZGRyICsgMCk7XG5cdFx0XHRcdGNvbnN0IGxlbiA9IGdldEludDY0KGFkZHIgKyA4KTtcblx0XHRcdFx0Y29uc3QgYSA9IG5ldyBBcnJheShsZW4pO1xuXHRcdFx0XHRmb3IgKGxldCBpID0gMDsgaSA8IGxlbjsgaSsrKSB7XG5cdFx0XHRcdFx0YVtpXSA9IGxvYWRWYWx1ZShhcnJheSArIGkgKiA4KTtcblx0XHRcdFx0fVxuXHRcdFx0XHRyZXR1cm4gYTtcblx0XHRcdH1cblxuXHRcdFx0Y29uc3QgbG9hZFN0cmluZyA9IChhZGRyKSA9PiB7XG5cdFx0XHRcdGNvbnN0IHNhZGRyID0gZ2V0SW50NjQoYWRkciArIDApO1xuXHRcdFx0XHRjb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7XG5cdFx0XHRcdHJldHVybiBkZWNvZGVyLmRlY29kZShuZXcgRGF0YVZpZXcodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIHNhZGRyLCBsZW4pKTtcblx0XHRcdH1cblxuXHRcdFx0Y29uc3QgdGltZU9yaWdpbiA9IERhdGUubm93KCkgLSBwZXJmb3JtYW5jZS5ub3coKTtcblx0XHRcdHRoaXMuaW1wb3J0T2JqZWN0ID0ge1xuXHRcdFx0XHRfZ290ZXN0OiB7XG5cdFx0XHRcdFx0YWRkOiAoYSwgYikgPT4gYSArIGIsXG5cdFx0XHRcdH0sXG5cdFx0XHRcdGdvanM6IHtcblx0XHRcdFx0XHQvLyBHbydzIFNQIGRvZXMgbm90IGNoYW5nZSBhcyBsb25nIGFzIG5vIEdvIGNvZGUgaXMgcnVubmluZy4gU29tZSBvcGVyYXRpb25zIChlLmcuIGNhbGxzLCBnZXR0ZXJzIGFuZCBzZXR0ZXJzKVxuXHRcdFx0XHRcdC8vIG1heSBzeW5jaHJvbm91c2x5IHRyaWdnZXIgYSBHbyBldmVudCBoYW5kbGVyLiBUaGlzIG1ha2VzIEdvIGNvZGUgZ2V0IGV4ZWN1dGVkIGluIHRoZSBtaWRkbGUgb2YgdGhlIGltcG9ydGVkXG5cdFx0XHRcdFx0Ly8gZnVuY3Rpb24uIEEgZ29yb3V0aW5lIGNhbiBzd2l0Y2ggdG8gYSBuZXcgc3RhY2sgaWYgdGhlIGN1cnJlbnQgc3RhY2sgaXMgdG9vIHNtYWxsIChzZWUgbW9yZXN0YWNrIGZ1bmN0aW9uKS5cblx0XHRcdFx0XHQvLyBUaGlzIGNoYW5nZXMgdGhlIFNQLCB0aHVzIHdlIGhhdmUgdG8gdXBkYXRlIHRoZSBTUCB1c2VkIGJ5IHRoZSBpbXBvcnRlZCBmdW5jdGlvbi5cblxuXHRcdFx0XHRcdC8vIGZ1bmMgd2FzbUV4aXQoY29kZSBpbnQzMilcblx0XHRcdFx0XHRcInJ1bnRpbWUud2FzbUV4aXRcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBjb2RlID0gdGhpcy5tZW0uZ2V0SW50MzIoc3AgKyA4LCB0cnVlKTtcblx0XHRcdFx0XHRcdHRoaXMuZXhpdGVkID0gdHJ1ZTtcblx0XHRcdFx0XHRcdGRlbGV0ZSB0aGlzLl9pbnN0O1xuXHRcdFx0XHRcdFx0ZGVsZXRlIHRoaXMuX3ZhbHVlcztcblx0XHRcdFx0XHRcdGRlbGV0ZSB0aGlzLl9nb1JlZkNvdW50cztcblx0XHRcdFx0XHRcdGRlbGV0ZSB0aGlzLl9pZHM7XG5cdFx0XHRcdFx0XHRkZWxldGUgdGhpcy5faWRQb29sO1xuXHRcdFx0XHRcdFx0dGhpcy5leGl0KGNvZGUpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHdhc21Xcml0ZShmZCB1aW50cHRyLCBwIHVuc2FmZS5Qb2ludGVyLCBuIGludDMyKVxuXHRcdFx0XHRcdFwicnVudGltZS53YXNtV3JpdGVcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBmZCA9IGdldEludDY0KHNwICsgOCk7XG5cdFx0XHRcdFx0XHRjb25zdCBwID0gZ2V0SW50NjQoc3AgKyAxNik7XG5cdFx0XHRcdFx0XHRjb25zdCBuID0gdGhpcy5tZW0uZ2V0SW50MzIoc3AgKyAyNCwgdHJ1ZSk7XG5cdFx0XHRcdFx0XHRmcy53cml0ZVN5bmMoZmQsIG5ldyBVaW50OEFycmF5KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLCBwLCBuKSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgcmVzZXRNZW1vcnlEYXRhVmlldygpXG5cdFx0XHRcdFx0XCJydW50aW1lLnJlc2V0TWVtb3J5RGF0YVZpZXdcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbSA9IG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlcik7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgbmFub3RpbWUxKCkgaW50NjRcblx0XHRcdFx0XHRcInJ1bnRpbWUubmFub3RpbWUxXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0c2V0SW50NjQoc3AgKyA4LCAodGltZU9yaWdpbiArIHBlcmZvcm1hbmNlLm5vdygpKSAqIDEwMDAwMDApO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHdhbGx0aW1lKCkgKHNlYyBpbnQ2NCwgbnNlYyBpbnQzMilcblx0XHRcdFx0XHRcInJ1bnRpbWUud2FsbHRpbWVcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBtc2VjID0gKG5ldyBEYXRlKS5nZXRUaW1lKCk7XG5cdFx0XHRcdFx0XHRzZXRJbnQ2NChzcCArIDgsIG1zZWMgLyAxMDAwKTtcblx0XHRcdFx0XHRcdHRoaXMubWVtLnNldEludDMyKHNwICsgMTYsIChtc2VjICUgMTAwMCkgKiAxMDAwMDAwLCB0cnVlKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyBzY2hlZHVsZVRpbWVvdXRFdmVudChkZWxheSBpbnQ2NCkgaW50MzJcblx0XHRcdFx0XHRcInJ1bnRpbWUuc2NoZWR1bGVUaW1lb3V0RXZlbnRcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBpZCA9IHRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRDtcblx0XHRcdFx0XHRcdHRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRCsrO1xuXHRcdFx0XHRcdFx0dGhpcy5fc2NoZWR1bGVkVGltZW91dHMuc2V0KGlkLCBzZXRUaW1lb3V0KFxuXHRcdFx0XHRcdFx0XHQoKSA9PiB7XG5cdFx0XHRcdFx0XHRcdFx0dGhpcy5fcmVzdW1lKCk7XG5cdFx0XHRcdFx0XHRcdFx0d2hpbGUgKHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLmhhcyhpZCkpIHtcblx0XHRcdFx0XHRcdFx0XHRcdC8vIGZvciBzb21lIHJlYXNvbiBHbyBmYWlsZWQgdG8gcmVnaXN0ZXIgdGhlIHRpbWVvdXQgZXZlbnQsIGxvZyBhbmQgdHJ5IGFnYWluXG5cdFx0XHRcdFx0XHRcdFx0XHQvLyAodGVtcG9yYXJ5IHdvcmthcm91bmQgZm9yIGh0dHBzOi8vZ2l0aHViLmNvbS9nb2xhbmcvZ28vaXNzdWVzLzI4OTc1KVxuXHRcdFx0XHRcdFx0XHRcdFx0Y29uc29sZS53YXJuKFwic2NoZWR1bGVUaW1lb3V0RXZlbnQ6IG1pc3NlZCB0aW1lb3V0IGV2ZW50XCIpO1xuXHRcdFx0XHRcdFx0XHRcdFx0dGhpcy5fcmVzdW1lKCk7XG5cdFx0XHRcdFx0XHRcdFx0fVxuXHRcdFx0XHRcdFx0XHR9LFxuXHRcdFx0XHRcdFx0XHRnZXRJbnQ2NChzcCArIDgpLFxuXHRcdFx0XHRcdFx0KSk7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRJbnQzMihzcCArIDE2LCBpZCwgdHJ1ZSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgY2xlYXJUaW1lb3V0RXZlbnQoaWQgaW50MzIpXG5cdFx0XHRcdFx0XCJydW50aW1lLmNsZWFyVGltZW91dEV2ZW50XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgaWQgPSB0aGlzLm1lbS5nZXRJbnQzMihzcCArIDgsIHRydWUpO1xuXHRcdFx0XHRcdFx0Y2xlYXJUaW1lb3V0KHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLmdldChpZCkpO1xuXHRcdFx0XHRcdFx0dGhpcy5fc2NoZWR1bGVkVGltZW91dHMuZGVsZXRlKGlkKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyBnZXRSYW5kb21EYXRhKHIgW11ieXRlKVxuXHRcdFx0XHRcdFwicnVudGltZS5nZXRSYW5kb21EYXRhXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y3J5cHRvLmdldFJhbmRvbVZhbHVlcyhsb2FkU2xpY2Uoc3AgKyA4KSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgZmluYWxpemVSZWYodiByZWYpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLmZpbmFsaXplUmVmXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgaWQgPSB0aGlzLm1lbS5nZXRVaW50MzIoc3AgKyA4LCB0cnVlKTtcblx0XHRcdFx0XHRcdHRoaXMuX2dvUmVmQ291bnRzW2lkXS0tO1xuXHRcdFx0XHRcdFx0aWYgKHRoaXMuX2dvUmVmQ291bnRzW2lkXSA9PT0gMCkge1xuXHRcdFx0XHRcdFx0XHRjb25zdCB2ID0gdGhpcy5fdmFsdWVzW2lkXTtcblx0XHRcdFx0XHRcdFx0dGhpcy5fdmFsdWVzW2lkXSA9IG51bGw7XG5cdFx0XHRcdFx0XHRcdHRoaXMuX2lkcy5kZWxldGUodik7XG5cdFx0XHRcdFx0XHRcdHRoaXMuX2lkUG9vbC5wdXNoKGlkKTtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyBzdHJpbmdWYWwodmFsdWUgc3RyaW5nKSByZWZcblx0XHRcdFx0XHRcInN5c2NhbGwvanMuc3RyaW5nVmFsXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDI0LCBsb2FkU3RyaW5nKHNwICsgOCkpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHZhbHVlR2V0KHYgcmVmLCBwIHN0cmluZykgcmVmXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlR2V0XCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgcmVzdWx0ID0gUmVmbGVjdC5nZXQobG9hZFZhbHVlKHNwICsgOCksIGxvYWRTdHJpbmcoc3AgKyAxNikpO1xuXHRcdFx0XHRcdFx0c3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmVcblx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyAzMiwgcmVzdWx0KTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZVNldCh2IHJlZiwgcCBzdHJpbmcsIHggcmVmKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZVNldFwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdFJlZmxlY3Quc2V0KGxvYWRWYWx1ZShzcCArIDgpLCBsb2FkU3RyaW5nKHNwICsgMTYpLCBsb2FkVmFsdWUoc3AgKyAzMikpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHZhbHVlRGVsZXRlKHYgcmVmLCBwIHN0cmluZylcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVEZWxldGVcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRSZWZsZWN0LmRlbGV0ZVByb3BlcnR5KGxvYWRWYWx1ZShzcCArIDgpLCBsb2FkU3RyaW5nKHNwICsgMTYpKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZUluZGV4KHYgcmVmLCBpIGludCkgcmVmXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlSW5kZXhcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgMjQsIFJlZmxlY3QuZ2V0KGxvYWRWYWx1ZShzcCArIDgpLCBnZXRJbnQ2NChzcCArIDE2KSkpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyB2YWx1ZVNldEluZGV4KHYgcmVmLCBpIGludCwgeCByZWYpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlU2V0SW5kZXhcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRSZWZsZWN0LnNldChsb2FkVmFsdWUoc3AgKyA4KSwgZ2V0SW50NjQoc3AgKyAxNiksIGxvYWRWYWx1ZShzcCArIDI0KSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgdmFsdWVDYWxsKHYgcmVmLCBtIHN0cmluZywgYXJncyBbXXJlZikgKHJlZiwgYm9vbClcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVDYWxsXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0dHJ5IHtcblx0XHRcdFx0XHRcdFx0Y29uc3QgdiA9IGxvYWRWYWx1ZShzcCArIDgpO1xuXHRcdFx0XHRcdFx0XHRjb25zdCBtID0gUmVmbGVjdC5nZXQodiwgbG9hZFN0cmluZyhzcCArIDE2KSk7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDMyKTtcblx0XHRcdFx0XHRcdFx0Y29uc3QgcmVzdWx0ID0gUmVmbGVjdC5hcHBseShtLCB2LCBhcmdzKTtcblx0XHRcdFx0XHRcdFx0c3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmVcblx0XHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDU2LCByZXN1bHQpO1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDY0LCAxKTtcblx0XHRcdFx0XHRcdH0gY2F0Y2ggKGVycikge1xuXHRcdFx0XHRcdFx0XHRzcCA9IHRoaXMuX2luc3QuZXhwb3J0cy5nZXRzcCgpID4+PiAwOyAvLyBzZWUgY29tbWVudCBhYm92ZVxuXHRcdFx0XHRcdFx0XHRzdG9yZVZhbHVlKHNwICsgNTYsIGVycik7XG5cdFx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgNjQsIDApO1xuXHRcdFx0XHRcdFx0fVxuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyBmdW5jIHZhbHVlSW52b2tlKHYgcmVmLCBhcmdzIFtdcmVmKSAocmVmLCBib29sKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUludm9rZVwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHRyeSB7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IHYgPSBsb2FkVmFsdWUoc3AgKyA4KTtcblx0XHRcdFx0XHRcdFx0Y29uc3QgYXJncyA9IGxvYWRTbGljZU9mVmFsdWVzKHNwICsgMTYpO1xuXHRcdFx0XHRcdFx0XHRjb25zdCByZXN1bHQgPSBSZWZsZWN0LmFwcGx5KHYsIHVuZGVmaW5lZCwgYXJncyk7XG5cdFx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyA0MCwgcmVzdWx0KTtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMSk7XG5cdFx0XHRcdFx0XHR9IGNhdGNoIChlcnIpIHtcblx0XHRcdFx0XHRcdFx0c3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmVcblx0XHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDQwLCBlcnIpO1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZU5ldyh2IHJlZiwgYXJncyBbXXJlZikgKHJlZiwgYm9vbClcblx0XHRcdFx0XHRcInN5c2NhbGwvanMudmFsdWVOZXdcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHR0cnkge1xuXHRcdFx0XHRcdFx0XHRjb25zdCB2ID0gbG9hZFZhbHVlKHNwICsgOCk7XG5cdFx0XHRcdFx0XHRcdGNvbnN0IGFyZ3MgPSBsb2FkU2xpY2VPZlZhbHVlcyhzcCArIDE2KTtcblx0XHRcdFx0XHRcdFx0Y29uc3QgcmVzdWx0ID0gUmVmbGVjdC5jb25zdHJ1Y3QodiwgYXJncyk7XG5cdFx0XHRcdFx0XHRcdHNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlXG5cdFx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyA0MCwgcmVzdWx0KTtcblx0XHRcdFx0XHRcdFx0dGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMSk7XG5cdFx0XHRcdFx0XHR9IGNhdGNoIChlcnIpIHtcblx0XHRcdFx0XHRcdFx0c3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmVcblx0XHRcdFx0XHRcdFx0c3RvcmVWYWx1ZShzcCArIDQwLCBlcnIpO1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTtcblx0XHRcdFx0XHRcdH1cblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZUxlbmd0aCh2IHJlZikgaW50XG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlTGVuZ3RoXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0c2V0SW50NjQoc3AgKyAxNiwgcGFyc2VJbnQobG9hZFZhbHVlKHNwICsgOCkubGVuZ3RoKSk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIHZhbHVlUHJlcGFyZVN0cmluZyh2IHJlZikgKHJlZiwgaW50KVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZVByZXBhcmVTdHJpbmdcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBzdHIgPSBlbmNvZGVyLmVuY29kZShTdHJpbmcobG9hZFZhbHVlKHNwICsgOCkpKTtcblx0XHRcdFx0XHRcdHN0b3JlVmFsdWUoc3AgKyAxNiwgc3RyKTtcblx0XHRcdFx0XHRcdHNldEludDY0KHNwICsgMjQsIHN0ci5sZW5ndGgpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHQvLyB2YWx1ZUxvYWRTdHJpbmcodiByZWYsIGIgW11ieXRlKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy52YWx1ZUxvYWRTdHJpbmdcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBzdHIgPSBsb2FkVmFsdWUoc3AgKyA4KTtcblx0XHRcdFx0XHRcdGxvYWRTbGljZShzcCArIDE2KS5zZXQoc3RyKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyB2YWx1ZUluc3RhbmNlT2YodiByZWYsIHQgcmVmKSBib29sXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLnZhbHVlSW5zdGFuY2VPZlwiOiAoc3ApID0+IHtcblx0XHRcdFx0XHRcdHNwID4+Pj0gMDtcblx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgMjQsIChsb2FkVmFsdWUoc3AgKyA4KSBpbnN0YW5jZW9mIGxvYWRWYWx1ZShzcCArIDE2KSkgPyAxIDogMCk7XG5cdFx0XHRcdFx0fSxcblxuXHRcdFx0XHRcdC8vIGZ1bmMgY29weUJ5dGVzVG9Hbyhkc3QgW11ieXRlLCBzcmMgcmVmKSAoaW50LCBib29sKVxuXHRcdFx0XHRcdFwic3lzY2FsbC9qcy5jb3B5Qnl0ZXNUb0dvXCI6IChzcCkgPT4ge1xuXHRcdFx0XHRcdFx0c3AgPj4+PSAwO1xuXHRcdFx0XHRcdFx0Y29uc3QgZHN0ID0gbG9hZFNsaWNlKHNwICsgOCk7XG5cdFx0XHRcdFx0XHRjb25zdCBzcmMgPSBsb2FkVmFsdWUoc3AgKyAzMik7XG5cdFx0XHRcdFx0XHRpZiAoIShzcmMgaW5zdGFuY2VvZiBVaW50OEFycmF5IHx8IHNyYyBpbnN0YW5jZW9mIFVpbnQ4Q2xhbXBlZEFycmF5KSkge1xuXHRcdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTtcblx0XHRcdFx0XHRcdFx0cmV0dXJuO1xuXHRcdFx0XHRcdFx0fVxuXHRcdFx0XHRcdFx0Y29uc3QgdG9Db3B5ID0gc3JjLnN1YmFycmF5KDAsIGRzdC5sZW5ndGgpO1xuXHRcdFx0XHRcdFx0ZHN0LnNldCh0b0NvcHkpO1xuXHRcdFx0XHRcdFx0c2V0SW50NjQoc3AgKyA0MCwgdG9Db3B5Lmxlbmd0aCk7XG5cdFx0XHRcdFx0XHR0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAxKTtcblx0XHRcdFx0XHR9LFxuXG5cdFx0XHRcdFx0Ly8gZnVuYyBjb3B5Qnl0ZXNUb0pTKGRzdCByZWYsIHNyYyBbXWJ5dGUpIChpbnQsIGJvb2wpXG5cdFx0XHRcdFx0XCJzeXNjYWxsL2pzLmNvcHlCeXRlc1RvSlNcIjogKHNwKSA9PiB7XG5cdFx0XHRcdFx0XHRzcCA+Pj49IDA7XG5cdFx0XHRcdFx0XHRjb25zdCBkc3QgPSBsb2FkVmFsdWUoc3AgKyA4KTtcblx0XHRcdFx0XHRcdGNvbnN0IHNyYyA9IGxvYWRTbGljZShzcCArIDE2KTtcblx0XHRcdFx0XHRcdGlmICghKGRzdCBpbnN0YW5jZW9mIFVpbnQ4QXJyYXkgfHwgZHN0IGluc3RhbmNlb2YgVWludDhDbGFtcGVkQXJyYXkpKSB7XG5cdFx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDApO1xuXHRcdFx0XHRcdFx0XHRyZXR1cm47XG5cdFx0XHRcdFx0XHR9XG5cdFx0XHRcdFx0XHRjb25zdCB0b0NvcHkgPSBzcmMuc3ViYXJyYXkoMCwgZHN0Lmxlbmd0aCk7XG5cdFx0XHRcdFx0XHRkc3Quc2V0KHRvQ29weSk7XG5cdFx0XHRcdFx0XHRzZXRJbnQ2NChzcCArIDQwLCB0b0NvcHkubGVuZ3RoKTtcblx0XHRcdFx0XHRcdHRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDEpO1xuXHRcdFx0XHRcdH0sXG5cblx0XHRcdFx0XHRcImRlYnVnXCI6ICh2YWx1ZSkgPT4ge1xuXHRcdFx0XHRcdFx0Y29uc29sZS5sb2codmFsdWUpO1xuXHRcdFx0XHRcdH0sXG5cdFx0XHRcdH1cblx0XHRcdH07XG5cdFx0fVxuXG5cdFx0YXN5bmMgcnVuKGluc3RhbmNlKSB7XG5cdFx0XHRpZiAoIShpbnN0YW5jZSBpbnN0YW5jZW9mIFdlYkFzc2VtYmx5Lkluc3RhbmNlKSkge1xuXHRcdFx0XHR0aHJvdyBuZXcgRXJyb3IoXCJHby5ydW46IFdlYkFzc2VtYmx5Lkluc3RhbmNlIGV4cGVjdGVkXCIpO1xuXHRcdFx0fVxuXHRcdFx0dGhpcy5faW5zdCA9IGluc3RhbmNlO1xuXHRcdFx0dGhpcy5tZW0gPSBuZXcgRGF0YVZpZXcodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIpO1xuXHRcdFx0dGhpcy5fdmFsdWVzID0gWyAvLyBKUyB2YWx1ZXMgdGhhdCBHbyBjdXJyZW50bHkgaGFzIHJlZmVyZW5jZXMgdG8sIGluZGV4ZWQgYnkgcmVmZXJlbmNlIGlkXG5cdFx0XHRcdE5hTixcblx0XHRcdFx0MCxcblx0XHRcdFx0bnVsbCxcblx0XHRcdFx0dHJ1ZSxcblx0XHRcdFx0ZmFsc2UsXG5cdFx0XHRcdGdsb2JhbFRoaXMsXG5cdFx0XHRcdHRoaXMsXG5cdFx0XHRdO1xuXHRcdFx0dGhpcy5fZ29SZWZDb3VudHMgPSBuZXcgQXJyYXkodGhpcy5fdmFsdWVzLmxlbmd0aCkuZmlsbChJbmZpbml0eSk7IC8vIG51bWJlciBvZiByZWZlcmVuY2VzIHRoYXQgR28gaGFzIHRvIGEgSlMgdmFsdWUsIGluZGV4ZWQgYnkgcmVmZXJlbmNlIGlkXG5cdFx0XHR0aGlzLl9pZHMgPSBuZXcgTWFwKFsgLy8gbWFwcGluZyBmcm9tIEpTIHZhbHVlcyB0byByZWZlcmVuY2UgaWRzXG5cdFx0XHRcdFswLCAxXSxcblx0XHRcdFx0W251bGwsIDJdLFxuXHRcdFx0XHRbdHJ1ZSwgM10sXG5cdFx0XHRcdFtmYWxzZSwgNF0sXG5cdFx0XHRcdFtnbG9iYWxUaGlzLCA1XSxcblx0XHRcdFx0W3RoaXMsIDZdLFxuXHRcdFx0XSk7XG5cdFx0XHR0aGlzLl9pZFBvb2wgPSBbXTsgICAvLyB1bnVzZWQgaWRzIHRoYXQgaGF2ZSBiZWVuIGdhcmJhZ2UgY29sbGVjdGVkXG5cdFx0XHR0aGlzLmV4aXRlZCA9IGZhbHNlOyAvLyB3aGV0aGVyIHRoZSBHbyBwcm9ncmFtIGhhcyBleGl0ZWRcblxuXHRcdFx0Ly8gUGFzcyBjb21tYW5kIGxpbmUgYXJndW1lbnRzIGFuZCBlbnZpcm9ubWVudCB2YXJpYWJsZXMgdG8gV2ViQXNzZW1ibHkgYnkgd3JpdGluZyB0aGVtIHRvIHRoZSBsaW5lYXIgbWVtb3J5LlxuXHRcdFx0bGV0IG9mZnNldCA9IDQwOTY7XG5cblx0XHRcdGNvbnN0IHN0clB0ciA9IChzdHIpID0+IHtcblx0XHRcdFx0Y29uc3QgcHRyID0gb2Zmc2V0O1xuXHRcdFx0XHRjb25zdCBieXRlcyA9IGVuY29kZXIuZW5jb2RlKHN0ciArIFwiXFwwXCIpO1xuXHRcdFx0XHRuZXcgVWludDhBcnJheSh0aGlzLm1lbS5idWZmZXIsIG9mZnNldCwgYnl0ZXMubGVuZ3RoKS5zZXQoYnl0ZXMpO1xuXHRcdFx0XHRvZmZzZXQgKz0gYnl0ZXMubGVuZ3RoO1xuXHRcdFx0XHRpZiAob2Zmc2V0ICUgOCAhPT0gMCkge1xuXHRcdFx0XHRcdG9mZnNldCArPSA4IC0gKG9mZnNldCAlIDgpO1xuXHRcdFx0XHR9XG5cdFx0XHRcdHJldHVybiBwdHI7XG5cdFx0XHR9O1xuXG5cdFx0XHRjb25zdCBhcmdjID0gdGhpcy5hcmd2Lmxlbmd0aDtcblxuXHRcdFx0Y29uc3QgYXJndlB0cnMgPSBbXTtcblx0XHRcdHRoaXMuYXJndi5mb3JFYWNoKChhcmcpID0+IHtcblx0XHRcdFx0YXJndlB0cnMucHVzaChzdHJQdHIoYXJnKSk7XG5cdFx0XHR9KTtcblx0XHRcdGFyZ3ZQdHJzLnB1c2goMCk7XG5cblx0XHRcdGNvbnN0IGtleXMgPSBPYmplY3Qua2V5cyh0aGlzLmVudikuc29ydCgpO1xuXHRcdFx0a2V5cy5mb3JFYWNoKChrZXkpID0+IHtcblx0XHRcdFx0YXJndlB0cnMucHVzaChzdHJQdHIoYCR7a2V5fT0ke3RoaXMuZW52W2tleV19YCkpO1xuXHRcdFx0fSk7XG5cdFx0XHRhcmd2UHRycy5wdXNoKDApO1xuXG5cdFx0XHRjb25zdCBhcmd2ID0gb2Zmc2V0O1xuXHRcdFx0YXJndlB0cnMuZm9yRWFjaCgocHRyKSA9PiB7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihvZmZzZXQsIHB0ciwgdHJ1ZSk7XG5cdFx0XHRcdHRoaXMubWVtLnNldFVpbnQzMihvZmZzZXQgKyA0LCAwLCB0cnVlKTtcblx0XHRcdFx0b2Zmc2V0ICs9IDg7XG5cdFx0XHR9KTtcblxuXHRcdFx0Ly8gVGhlIGxpbmtlciBndWFyYW50ZWVzIGdsb2JhbCBkYXRhIHN0YXJ0cyBmcm9tIGF0IGxlYXN0IHdhc21NaW5EYXRhQWRkci5cblx0XHRcdC8vIEtlZXAgaW4gc3luYyB3aXRoIGNtZC9saW5rL2ludGVybmFsL2xkL2RhdGEuZ286d2FzbU1pbkRhdGFBZGRyLlxuXHRcdFx0Y29uc3Qgd2FzbU1pbkRhdGFBZGRyID0gNDA5NiArIDgxOTI7XG5cdFx0XHRpZiAob2Zmc2V0ID49IHdhc21NaW5EYXRhQWRkcikge1xuXHRcdFx0XHR0aHJvdyBuZXcgRXJyb3IoXCJ0b3RhbCBsZW5ndGggb2YgY29tbWFuZCBsaW5lIGFuZCBlbnZpcm9ubWVudCB2YXJpYWJsZXMgZXhjZWVkcyBsaW1pdFwiKTtcblx0XHRcdH1cblxuXHRcdFx0dGhpcy5faW5zdC5leHBvcnRzLnJ1bihhcmdjLCBhcmd2KTtcblx0XHRcdGlmICh0aGlzLmV4aXRlZCkge1xuXHRcdFx0XHR0aGlzLl9yZXNvbHZlRXhpdFByb21pc2UoKTtcblx0XHRcdH1cblx0XHRcdGF3YWl0IHRoaXMuX2V4aXRQcm9taXNlO1xuXHRcdH1cblxuXHRcdF9yZXN1bWUoKSB7XG5cdFx0XHRpZiAodGhpcy5leGl0ZWQpIHtcblx0XHRcdFx0dGhyb3cgbmV3IEVycm9yKFwiR28gcHJvZ3JhbSBoYXMgYWxyZWFkeSBleGl0ZWRcIik7XG5cdFx0XHR9XG5cdFx0XHR0aGlzLl9pbnN0LmV4cG9ydHMucmVzdW1lKCk7XG5cdFx0XHRpZiAodGhpcy5leGl0ZWQpIHtcblx0XHRcdFx0dGhpcy5fcmVzb2x2ZUV4aXRQcm9taXNlKCk7XG5cdFx0XHR9XG5cdFx0fVxuXG5cdFx0X21ha2VGdW5jV3JhcHBlcihpZCkge1xuXHRcdFx0Y29uc3QgZ28gPSB0aGlzO1xuXHRcdFx0cmV0dXJuIGZ1bmN0aW9uICgpIHtcblx0XHRcdFx0Y29uc3QgZXZlbnQgPSB7IGlkOiBpZCwgdGhpczogdGhpcywgYXJnczogYXJndW1lbnRzIH07XG5cdFx0XHRcdGdvLl9wZW5kaW5nRXZlbnQgPSBldmVudDtcblx0XHRcdFx0Z28uX3Jlc3VtZSgpO1xuXHRcdFx0XHRyZXR1cm4gZXZlbnQucmVzdWx0O1xuXHRcdFx0fTtcblx0XHR9XG5cdH1cbn0pKCk7XG5cblxuXG4vLyBXRUJQQUNLIEZPT1RFUiAvL1xuLy8gLi9zcmMvdmlld3Mvd2FzbV9leGVjLmpzIl0sInNvdXJjZVJvb3QiOiIifQ==