#!/usr/bin/env python3
"""Regroupe des modèles glTF/GLB (packs CC0 Kenney / Quaternius) en un seul GLB statique.
Chaque modèle devient un nœud racine nommé par sa clé. Textures dédoublonnées, animations/squelettes retirés.
Usage : citypack.py liste.txt sortie.glb   (liste : « clé chemin » par ligne)"""
import sys, os, json, struct, base64, hashlib, io
import numpy as np
QUANT = os.environ.get('QUANT', '1') == '1'
CT = {5120: np.int8, 5121: np.uint8, 5122: np.int16, 5123: np.uint16, 5125: np.uint32, 5126: np.float32}
NC = {'SCALAR': 1, 'VEC2': 2, 'VEC3': 3, 'VEC4': 4}
KEEP_ATTR = ('POSITION', 'NORMAL', 'TEXCOORD_0', 'COLOR_0')
out = {'asset': {'version': '2.0', 'generator': 'velocite citypack'}, 'scene': 0, 'scenes': [{'nodes': []}],
       'nodes': [], 'meshes': [], 'materials': [], 'textures': [], 'images': [], 'samplers': [], 'accessors': [], 'bufferViews': [], 'buffers': []}
blob = bytearray(); img_seen = {}
def add_view(data, target=None, stride=None):
    while len(blob) % 4: blob.append(0)
    v = {'buffer': 0, 'byteOffset': len(blob), 'byteLength': len(data)}
    if target: v['target'] = target
    if stride: v['byteStride'] = stride
    blob.extend(data); out['bufferViews'].append(v); return len(out['bufferViews']) - 1
def load(path):
    raw = open(path, 'rb').read(); base = os.path.dirname(path)
    if raw[:4] == b'glTF':
        L = struct.unpack('<I', raw[12:16])[0]; j = json.loads(raw[20:20 + L]); o = 20 + L
        bins = [raw[o + 8:o + 8 + struct.unpack('<I', raw[o:o + 4])[0]]] if o < len(raw) else []
    else:
        j = json.loads(raw); bins = []
        for b in j.get('buffers', []):
            u = b['uri']
            bins.append(base64.b64decode(u.split(',', 1)[1]) if u.startswith('data:') else open(os.path.join(base, u), 'rb').read())
    return j, bins, base
def view_bytes(j, bins, vi):
    v = j['bufferViews'][vi]; return bins[v.get('buffer', 0)][v.get('byteOffset', 0):v.get('byteOffset', 0) + v['byteLength']]
def read_acc(j, bins, ai):
    a = j['accessors'][ai]; v = j['bufferViews'][a['bufferView']]; dt = np.dtype(CT[a['componentType']]); n = NC[a['type']]
    buf = bins[v.get('buffer', 0)]; off = v.get('byteOffset', 0) + a.get('byteOffset', 0); st = v.get('byteStride', dt.itemsize * n)
    arr = np.ndarray((a['count'], n), dtype=dt, buffer=buf, offset=off, strides=(st, dt.itemsize)).astype(np.float64)
    if a.get('normalized'):
        arr = np.maximum(arr / {np.int8: 127, np.uint8: 255, np.int16: 32767, np.uint16: 65535}[CT[a['componentType']]], -1)
    return arr
def new_acc(arr, ctype, typ, normalized=False, minmax=False, stride=None, target=34962):
    data = arr.tobytes(); a = {'bufferView': add_view(data, target, stride), 'componentType': ctype, 'count': int(arr.shape[0]), 'type': typ}
    if normalized: a['normalized'] = True
    if minmax: a['min'] = [float(x) for x in arr.min(0)[:NC[typ]]]; a['max'] = [float(x) for x in arr.max(0)[:NC[typ]]]
    out['accessors'].append(a); return len(out['accessors']) - 1
def pack(key, path):
    j, bins, base = load(path)
    amap, vmap, tmap, mmap, imap, smap = {}, {}, {}, {}, {}, {}
    def acc(ai):
        if ai in amap: return amap[ai]
        a = dict(j['accessors'][ai]); vi = a['bufferView']
        if vi not in vmap:
            v = j['bufferViews'][vi]; vmap[vi] = add_view(view_bytes(j, bins, vi), v.get('target'), v.get('byteStride'))
        a['bufferView'] = vmap[vi]; a.pop('sparse', None)
        out['accessors'].append(a); amap[ai] = len(out['accessors']) - 1; return amap[ai]
    def image(ii):
        if ii in imap: return imap[ii]
        im = j['images'][ii]
        if 'uri' in im:
            u = im['uri']; data = base64.b64decode(u.split(',', 1)[1]) if u.startswith('data:') else open(os.path.join(base, u.replace('%20', ' ')), 'rb').read()
            mime = 'image/jpeg' if u.lower().endswith(('.jpg', '.jpeg')) else 'image/png'
        else: data = view_bytes(j, bins, im['bufferView']); mime = im.get('mimeType', 'image/png')
        if len(data) > 200000:
            from PIL import Image
            im2 = Image.open(io.BytesIO(data)).convert('RGB'); im2.thumbnail((512, 512)); b = io.BytesIO(); im2.save(b, 'JPEG', quality=82); data = b.getvalue(); mime = 'image/jpeg'
        h = hashlib.sha1(data).hexdigest()
        if h not in img_seen:
            out['images'].append({'bufferView': add_view(data), 'mimeType': mime, 'name': os.path.basename(im.get('uri', key))}); img_seen[h] = len(out['images']) - 1
        imap[ii] = img_seen[h]; return imap[ii]
    def tex(ti):
        if ti in tmap: return tmap[ti]
        t = dict(j['textures'][ti]); src = t.get('source')
        nt = {'source': image(src)} if src is not None else {}
        if 'sampler' in t:
            s = json.dumps(j['samplers'][t['sampler']], sort_keys=True)
            if s not in smap: out['samplers'].append(json.loads(s)); smap[s] = len(out['samplers']) - 1
            nt['sampler'] = smap[s]
        # même image + même sampler = même texture
        for k, e in enumerate(out['textures']):
            if e == nt: tmap[ti] = k; return k
        out['textures'].append(nt); tmap[ti] = len(out['textures']) - 1; return tmap[ti]
    def mat(mi):
        if mi in mmap: return mmap[mi]
        m = json.loads(json.dumps(j['materials'][mi])); m.pop('extensions', None)
        def fix(o):
            if isinstance(o, dict):
                for k, v in list(o.items()):
                    if k in ('baseColorTexture', 'normalTexture', 'metallicRoughnessTexture', 'emissiveTexture', 'occlusionTexture') and isinstance(v, dict):
                        if k in ('baseColorTexture',): v['index'] = tex(v['index']); v.pop('texCoord', None); v.pop('extensions', None)
                        else: o.pop(k)
                    else: fix(v)
        fix(m); m['name'] = key + ':' + m.get('name', str(mi))
        out['materials'].append(m); mmap[mi] = len(out['materials']) - 1; return mmap[mi]
    nmap = {}
    def node(ni):
        prims = None; qn = None
        n = j['nodes'][ni]; o = {k: n[k] for k in ('translation', 'rotation', 'scale', 'matrix', 'name') if k in n}
        if 'mesh' in n:
            me = j['meshes'][n['mesh']]; prims = []
            pl = [p for p in me['primitives'] if p.get('mode', 4) == 4]
            qn = None
            if QUANT and pl:
                P = [read_acc(j, bins, p['attributes']['POSITION']) for p in pl]
                lo = np.min([x.min(0) for x in P], 0); hi = np.max([x.max(0) for x in P], 0); c = (lo + hi) / 2; e = max(float((hi - lo).max()) / 2, 1e-6)
                qn = {'translation': [float(x) for x in c], 'scale': [e, e, e]}
            for k, p in enumerate(pl):
                if qn is not None:
                    attrs = {}
                    pos = np.clip(np.round((P[k] - c) / e * 32767), -32767, 32767).astype(np.int16)
                    pos4 = np.zeros((len(pos), 4), np.int16); pos4[:, :3] = pos
                    attrs['POSITION'] = new_acc(pos4, 5122, 'VEC3', True, True, 8)
                    out['accessors'][-1]['min'] = [int(x) for x in pos.min(0)]; out['accessors'][-1]['max'] = [int(x) for x in pos.max(0)]
                    if 'NORMAL' in p['attributes']:
                        nn = read_acc(j, bins, p['attributes']['NORMAL']); nn /= np.maximum(np.linalg.norm(nn, axis=1, keepdims=True), 1e-9)
                        n4 = np.zeros((len(nn), 4), np.int8); n4[:, :3] = np.round(nn * 127).astype(np.int8)
                        attrs['NORMAL'] = new_acc(n4, 5120, 'VEC3', True, False, 4)
                    if 'TEXCOORD_0' in p['attributes']:
                        uv = read_acc(j, bins, p['attributes']['TEXCOORD_0'])
                        if uv.min() >= 0 and uv.max() <= 1: attrs['TEXCOORD_0'] = new_acc(np.round(uv * 65535).astype(np.uint16), 5123, 'VEC2', True)
                        else: attrs['TEXCOORD_0'] = new_acc(uv.astype(np.float32), 5126, 'VEC2')
                    if 'COLOR_0' in p['attributes']: attrs['COLOR_0'] = acc(p['attributes']['COLOR_0'])
                    q = {'attributes': attrs}
                else:
                    q = {'attributes': {a: acc(i) for a, i in p['attributes'].items() if a in KEEP_ATTR}}
                if 'indices' in p:
                    if QUANT and j['accessors'][p['indices']]['componentType'] == 5125 and j['accessors'][p['indices']]['count'] and read_acc(j, bins, p['indices']).max() < 65535:
                        q['indices'] = new_acc(read_acc(j, bins, p['indices']).astype(np.uint16).reshape(-1), 5123, 'SCALAR', target=34963)
                    else: q['indices'] = acc(p['indices'])
                if 'material' in p: q['material'] = mat(p['material'])
                prims.append(q)
            if prims:
                out['meshes'].append({'primitives': prims})
                if qn is not None: qn['mesh'] = len(out['meshes']) - 1; out['nodes'].append(qn); qkid = len(out['nodes']) - 1
                else: o['mesh'] = len(out['meshes']) - 1
        kids = [node(c) for c in n.get('children', [])]
        if 'mesh' in n and qn is not None and prims: kids.insert(0, qkid)
        if kids: o['children'] = kids
        out['nodes'].append(o); return len(out['nodes']) - 1
    roots = [node(r) for r in j['scenes'][j.get('scene', 0)]['nodes']]
    out['nodes'].append({'name': key, 'children': roots}); out['scenes'][0]['nodes'].append(len(out['nodes']) - 1)
lst, dst = sys.argv[1], sys.argv[2]
for line in open(lst):
    line = line.strip()
    if not line or line.startswith('#'): continue
    key, path = line.split(None, 1); pack(key, path)
while len(blob) % 4: blob.append(0)
out['buffers'] = [{'byteLength': len(blob)}]
for k in ('textures', 'images', 'samplers'):
    if not out[k]: out.pop(k)
if QUANT: out['extensionsUsed'] = out['extensionsRequired'] = ['KHR_mesh_quantization']
js = json.dumps(out, separators=(',', ':')).encode()
while len(js) % 4: js += b' '
with open(dst, 'wb') as f:
    f.write(struct.pack('<III', 0x46546C67, 2, 12 + 8 + len(js) + 8 + len(blob)))
    f.write(struct.pack('<II', len(js), 0x4E4F534A)); f.write(js)
    f.write(struct.pack('<II', len(blob), 0x004E4942)); f.write(blob)
print(dst, len(out['scenes'][0]['nodes']), 'modèles', round((len(js) + len(blob)) / 1048576, 2), 'Mo', len(out.get('images', [])), 'images', len(out['materials']), 'matériaux')
