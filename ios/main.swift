// Vélocité pour iPhone / iPad : une fenêtre WKWebView plein écran qui sert les fichiers du jeu rangés dans l'appli.
// Les fichiers sont servis par une adresse interne velocite://localhost/ (comme l'appli Android avec appassets) :
// pas de serveur, pas de réseau nécessaire pour jouer en solo ; stockage, WebGL, micro et modèles 3D fonctionnent.
// Construite sans signature par GitHub Actions (.github/workflows/ipa.yml) : LiveContainer, SideStore, AltStore… la signent eux-mêmes.
import UIKit
import WebKit
import AVFoundation
import UniformTypeIdentifiers

final class GameFiles: NSObject, WKURLSchemeHandler {
    let root: URL
    init(root: URL) { self.root = root.standardizedFileURL }

    func webView(_ webView: WKWebView, start task: WKURLSchemeTask) {
        guard let url = task.request.url else { task.didFailWithError(URLError(.badURL)); return }
        var path = url.path.removingPercentEncoding ?? url.path
        if path.isEmpty || path == "/" { path = "/index.html" }
        let file = root.appendingPathComponent(String(path.dropFirst())).standardizedFileURL
        guard file.path.hasPrefix(root.path + "/"),
              let data = try? Data(contentsOf: file, options: .mappedIfSafe) else {
            let r = HTTPURLResponse(url: url, statusCode: 404, httpVersion: "HTTP/1.1", headerFields: ["Content-Type": "text/plain"])!
            task.didReceive(r); task.didReceive(Data()); task.didFinish(); return
        }
        let ext = file.pathExtension.lowercased()
        let mime = ext == "glb" ? "model/gltf-binary" : ext == "js" ? "text/javascript" : ext == "wasm" ? "application/wasm"
            : UTType(filenameExtension: ext)?.preferredMIMEType ?? "application/octet-stream"
        var headers = ["Content-Type": mime, "Cache-Control": "no-cache", "Accept-Ranges": "bytes", "Access-Control-Allow-Origin": "*"]
        var status = 200, body = data
        // musique et sons : le lecteur audio demande des morceaux du fichier (Range)
        if let rg = task.request.value(forHTTPHeaderField: "Range"), rg.hasPrefix("bytes="), !data.isEmpty {
            let p = rg.dropFirst(6).split(separator: "-", omittingEmptySubsequences: false)
            let n = data.count
            var a = Int(p.first ?? "") ?? 0, b = p.count > 1 ? (Int(p[1]) ?? n - 1) : n - 1
            if p.first == "" , p.count > 1, let suf = Int(p[1]) { a = max(0, n - suf); b = n - 1 }
            b = min(b, n - 1)
            if a <= b {
                body = data.subdata(in: a..<(b + 1)); status = 206
                headers["Content-Range"] = "bytes \(a)-\(b)/\(n)"
            }
        }
        headers["Content-Length"] = String(body.count)
        task.didReceive(HTTPURLResponse(url: url, statusCode: status, httpVersion: "HTTP/1.1", headerFields: headers)!)
        task.didReceive(body)
        task.didFinish()
    }
    func webView(_ webView: WKWebView, stop task: WKURLSchemeTask) {}
}

final class GameVC: UIViewController, WKNavigationDelegate, WKUIDelegate, WKScriptMessageHandler {
    var web: WKWebView!
    static let home = URL(string: "velocite://localhost/index.html")!

    override func loadView() {
        let cfg = WKWebViewConfiguration()
        cfg.setURLSchemeHandler(GameFiles(root: Bundle.main.resourceURL!.appendingPathComponent("www")), forURLScheme: "velocite")
        cfg.allowsInlineMediaPlayback = true
        cfg.mediaTypesRequiringUserActionForPlayback = []
        cfg.preferences.javaScriptCanOpenWindowsAutomatically = true
        cfg.websiteDataStore = .default()
        // le jeu sait qu'il tourne dans l'appli iOS ; erreurs JavaScript renvoyées dans le journal de l'appli (pour le débogage)
        let js = """
        window.__VELOCITE_IOS = true;
        (function(){ var h = window.webkit && webkit.messageHandlers && webkit.messageHandlers.log; if (!h) return;
          window.addEventListener('error', function(e){ try { h.postMessage('ERR ' + e.message + ' @' + (e.filename||'') + ':' + (e.lineno||0)); } catch(_){} });
          window.addEventListener('unhandledrejection', function(e){ try { h.postMessage('REJ ' + String(e.reason && (e.reason.stack || e.reason))); } catch(_){} });
          function probe(tag){ try { var c = document.createElement('canvas'), ob = document.getElementById('obsBoot');
            fetch('version.json', { cache: 'no-store' }).then(function(r){ return r.json(); }).catch(function(){ return {}; }).then(function(j){
              h.postMessage('STATE ' + tag + ' version=' + (j.v || '?') + ' secure=' + window.isSecureContext +
                ' subtle=' + !!(window.crypto && crypto.subtle) + ' webgl=' + !!(c.getContext('webgl2') || c.getContext('webgl')) +
                ' storage=' + (function(){ try { localStorage.setItem('velocite_iosProbe', '1'); return localStorage.getItem('velocite_iosProbe') === '1'; } catch(_) { return false; } })() +
                ' three=' + (typeof THREE !== 'undefined') + ' size=' + innerWidth + 'x' + innerHeight +
                ' obsBoot=' + (ob ? 'visible bad=' + ob.querySelectorAll('.bad').length : 'gone') +
                ' text=' + (ob ? ob.innerText : (document.body.innerText || '')).replace(/\\s+/g, ' ').slice(0, 220));
            });
          } catch(e) { h.postMessage('STATE error ' + e); } }
          setTimeout(function(){ probe('t15'); }, 15000);
          setTimeout(function(){ probe('t40'); }, 40000);
        })();
        """
        cfg.userContentController.addUserScript(WKUserScript(source: js, injectionTime: .atDocumentStart, forMainFrameOnly: true))
        cfg.userContentController.add(self, name: "log")
        web = WKWebView(frame: .zero, configuration: cfg)
        web.isOpaque = false
        web.backgroundColor = UIColor(red: 0.055, green: 0.063, blue: 0.075, alpha: 1)
        web.scrollView.isScrollEnabled = false
        web.scrollView.bounces = false
        web.scrollView.contentInsetAdjustmentBehavior = .never
        web.allowsBackForwardNavigationGestures = false
        web.allowsLinkPreview = false
        if #available(iOS 16.4, *) { web.isInspectable = true }   // Safari sur Mac → Développement → inspecter
        web.navigationDelegate = self
        web.uiDelegate = self
        view = web
        web.load(URLRequest(url: GameVC.home))
    }

    override var prefersStatusBarHidden: Bool { true }
    override var prefersHomeIndicatorAutoHidden: Bool { true }
    override var preferredScreenEdgesDeferringSystemGestures: UIRectEdge { .all }
    override var supportedInterfaceOrientations: UIInterfaceOrientationMask { .landscape }

    func userContentController(_ c: WKUserContentController, didReceive m: WKScriptMessage) {
        NSLog("[Velocite] %@", String(describing: m.body))
    }

    // liens extérieurs (Discord, GitHub…) : ouverts dans Safari, le jeu reste dans l'appli
    func webView(_ w: WKWebView, decidePolicyFor a: WKNavigationAction, decisionHandler: @escaping (WKNavigationActionPolicy) -> Void) {
        if let u = a.request.url, let sc = u.scheme?.lowercased(), !["velocite", "about", "blob", "data", "javascript"].contains(sc),
           a.targetFrame?.isMainFrame ?? true {
            UIApplication.shared.open(u); decisionHandler(.cancel); return
        }
        decisionHandler(.allow)
    }
    func webView(_ w: WKWebView, createWebViewWith c: WKWebViewConfiguration, for a: WKNavigationAction, windowFeatures f: WKWindowFeatures) -> WKWebView? {
        if let u = a.request.url { UIApplication.shared.open(u) }
        return nil
    }
    // la mémoire du téléphone est pleine et iOS a fermé la page : on la relance
    func webViewWebContentProcessDidTerminate(_ w: WKWebView) { w.load(URLRequest(url: GameVC.home)) }

    // micro (chat vocal) : iOS demande l'autorisation une fois, la page n'a pas besoin de redemander
    @available(iOS 15.0, *)
    func webView(_ w: WKWebView, requestMediaCapturePermissionFor o: WKSecurityOrigin, initiatedByFrame f: WKFrameInfo, type: WKMediaCaptureType, decisionHandler: @escaping (WKPermissionDecision) -> Void) {
        decisionHandler(o.protocol == "velocite" ? .grant : .prompt)
    }

    // alert / confirm / prompt de la page
    func webView(_ w: WKWebView, runJavaScriptAlertPanelWithMessage msg: String, initiatedByFrame f: WKFrameInfo, completionHandler: @escaping () -> Void) {
        let a = UIAlertController(title: nil, message: msg, preferredStyle: .alert)
        a.addAction(UIAlertAction(title: "OK", style: .default) { _ in completionHandler() })
        present(a, animated: true)
    }
    func webView(_ w: WKWebView, runJavaScriptConfirmPanelWithMessage msg: String, initiatedByFrame f: WKFrameInfo, completionHandler: @escaping (Bool) -> Void) {
        let a = UIAlertController(title: nil, message: msg, preferredStyle: .alert)
        a.addAction(UIAlertAction(title: "Annuler", style: .cancel) { _ in completionHandler(false) })
        a.addAction(UIAlertAction(title: "OK", style: .default) { _ in completionHandler(true) })
        present(a, animated: true)
    }
    func webView(_ w: WKWebView, runJavaScriptTextInputPanelWithPrompt p: String, defaultText d: String?, initiatedByFrame f: WKFrameInfo, completionHandler: @escaping (String?) -> Void) {
        let a = UIAlertController(title: nil, message: p, preferredStyle: .alert)
        a.addTextField { $0.text = d }
        a.addAction(UIAlertAction(title: "Annuler", style: .cancel) { _ in completionHandler(nil) })
        a.addAction(UIAlertAction(title: "OK", style: .default) { _ in completionHandler(a.textFields?.first?.text) })
        present(a, animated: true)
    }
}

@main
final class AppDelegate: UIResponder, UIApplicationDelegate {
    var window: UIWindow?
    func application(_ app: UIApplication, didFinishLaunchingWithOptions o: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        // le son du jeu joue même avec le bouton silencieux, sans couper la musique des autres applis
        try? AVAudioSession.sharedInstance().setCategory(.playback, mode: .default, options: [.mixWithOthers])
        try? AVAudioSession.sharedInstance().setActive(true)
        UIApplication.shared.isIdleTimerDisabled = true   // l'écran ne s'éteint pas pendant une partie
        let w = UIWindow(frame: UIScreen.main.bounds)
        w.backgroundColor = .black
        w.rootViewController = GameVC()
        w.makeKeyAndVisible()
        window = w
        return true
    }
}
