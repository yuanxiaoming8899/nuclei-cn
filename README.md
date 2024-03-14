<div class="Box-sc-g0xbh4-0 bJMeLZ js-snippet-clipboard-copy-unpositioned" data-hpc="true"><article class="markdown-body entry-content container-lg" itemprop="text"><div class="markdown-heading" dir="auto"><h1 align="center" tabindex="-1" class="heading-element" dir="auto">
  <br>
  <a href="https://nuclei.projectdiscovery.io" rel="nofollow"><img src="/projectdiscovery/nuclei/raw/dev/static/nuclei-logo.png" width="200px" alt="细胞核" style="max-width: 100%;"></a>
</h1><a id="user-content-----" class="anchor" aria-label="永久链接：" href="#----"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<div class="markdown-heading" dir="auto"><h4 align="center" tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">基于简单的 YAML DSL 的快速且可定制的漏洞扫描器。</font></font></h4><a id="user-content-fast-and-customisable-vulnerability-scanner-based-on-simple-yaml-based-dsl" class="anchor" aria-label="永久链接：基于简单的 YAML DSL 的快速且可定制的漏洞扫描器。" href="#fast-and-customisable-vulnerability-scanner-based-on-simple-yaml-based-dsl"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p align="center" dir="auto">
<a target="_blank" rel="noopener noreferrer nofollow" href="https://camo.githubusercontent.com/d68fc253a0221c3a01c8be387d419427fba8e666a00546601856ec2d13ef1a3e/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f676f2d6d6f642f676f2d76657273696f6e2f70726f6a656374646973636f766572792f6e75636c6569"><img src="https://camo.githubusercontent.com/d68fc253a0221c3a01c8be387d419427fba8e666a00546601856ec2d13ef1a3e/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f676f2d6d6f642f676f2d76657273696f6e2f70726f6a656374646973636f766572792f6e75636c6569" data-canonical-src="https://img.shields.io/github/go-mod/go-version/projectdiscovery/nuclei" style="max-width: 100%;"></a>
<a href="https://github.com/projectdiscovery/nuclei/releases"><img src="https://camo.githubusercontent.com/b44ce2ea65d2ba298a902f4ee199602b31bd511032c71b0d02388f96448e3270/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f646f776e6c6f6164732f70726f6a656374646973636f766572792f6e75636c65692f746f74616c" data-canonical-src="https://img.shields.io/github/downloads/projectdiscovery/nuclei/total" style="max-width: 100%;">
</a><a href="https://github.com/projectdiscovery/nuclei/graphs/contributors"><img src="https://camo.githubusercontent.com/42bda7ab09e744c0ec515f95180b9daa850b8821d429483af09798febe492354/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f636f6e7472696275746f72732d616e6f6e2f70726f6a656374646973636f766572792f6e75636c6569" data-canonical-src="https://img.shields.io/github/contributors-anon/projectdiscovery/nuclei" style="max-width: 100%;">
</a><a href="https://github.com/projectdiscovery/nuclei/releases/"><img src="https://camo.githubusercontent.com/65facb7801ce8c23f2ef52e63a0638efad04f424a004df2bb696dbdcf58a7438/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f72656c656173652f70726f6a656374646973636f766572792f6e75636c6569" data-canonical-src="https://img.shields.io/github/release/projectdiscovery/nuclei" style="max-width: 100%;">
</a><a href="https://github.com/projectdiscovery/nuclei/issues"><img src="https://camo.githubusercontent.com/91c4ce02f611bfffabcbc320b74bf9171f670896b9198570a3c753c3b622a0c0/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f6973737565732d7261772f70726f6a656374646973636f766572792f6e75636c6569" data-canonical-src="https://img.shields.io/github/issues-raw/projectdiscovery/nuclei" style="max-width: 100%;">
</a><a href="https://github.com/projectdiscovery/nuclei/discussions"><img src="https://camo.githubusercontent.com/a27531114868583323654e00cc47c0a6a5622169df4167325f6252dae11a8901/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f64697363757373696f6e732f70726f6a656374646973636f766572792f6e75636c6569" data-canonical-src="https://img.shields.io/github/discussions/projectdiscovery/nuclei" style="max-width: 100%;">
</a><a href="https://discord.gg/projectdiscovery" rel="nofollow"><img src="https://camo.githubusercontent.com/cf5aeb9adbbb9a43aa5afcd14f2d6dec0f4d4661b226caa4640d9d9929e91b6a/68747470733a2f2f696d672e736869656c64732e696f2f646973636f72642f3639353634353233373431383133313530372e7376673f6c6f676f3d646973636f7264" data-canonical-src="https://img.shields.io/discord/695645237418131507.svg?logo=discord" style="max-width: 100%;"></a>
<a href="https://twitter.com/pdnuclei" rel="nofollow"><img src="https://camo.githubusercontent.com/9c1065af7ffbccf16141d6c039222675a25df94c74aa7a4d04eeb2508df2bd35/68747470733a2f2f696d672e736869656c64732e696f2f747769747465722f666f6c6c6f772f70646e75636c65692e7376673f6c6f676f3d74776974746572" data-canonical-src="https://img.shields.io/twitter/follow/pdnuclei.svg?logo=twitter" style="max-width: 100%;"></a>
</p>
<p align="center" dir="auto">
  <a href="#how-it-works"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">如何</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="#install-nuclei"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">安装</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://docs.projectdiscovery.io/tools/nuclei/" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">文档</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="#credits"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">制作人员</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://nuclei.projectdiscovery.io/faq/nuclei/" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">常见问题解答</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://discord.gg/projectdiscovery" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">加入 Discord</font></font></a>
</p>
<p align="center" dir="auto">
  <a href="https://github.com/projectdiscovery/nuclei/blob/main/README.md"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">英语</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://github.com/projectdiscovery/nuclei/blob/main/README_CN.md"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">中文</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://github.com/projectdiscovery/nuclei/blob/main/README_KR.md"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">韩语</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">•
  </font></font><a href="https://github.com/projectdiscovery/nuclei/blob/main/README_ID.md"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">印度尼西亚</font></font></a>
</p>
<hr>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Nuclei 用于基于模板跨目标发送请求，从而实现零误报并提供对大量主机的快速扫描。</font><font style="vertical-align: inherit;">Nuclei 提供对各种协议的扫描，包括 TCP、DNS、HTTP、SSL、File、Whois、Websocket、Headless、Code 等。凭借强大而灵活的模板，Nuclei 可用于对各种安全检查进行建模。</font></font></p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">我们有一个</font></font><a href="https://github.com/projectdiscovery/nuclei-templates"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">专门的存储库</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">，其中包含由</font></font><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">300 多名</font></font></strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">安全研究人员和工程师贡献的各种类型的漏洞模板。</font></font></p>
<div class="markdown-heading" dir="auto"><h2 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">怎么运行的</font></font></h2><a id="user-content-how-it-works" class="anchor" aria-label="永久链接：它是如何工作的" href="#how-it-works"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<div class="markdown-heading" dir="auto"><h3 align="center" tabindex="-1" class="heading-element" dir="auto">
  <a target="_blank" rel="noopener noreferrer" href="/projectdiscovery/nuclei/blob/dev/static/nuclei-flow.jpg"><img src="/projectdiscovery/nuclei/raw/dev/static/nuclei-flow.jpg" alt="核流" width="700px" style="max-width: 100%;"></a>
</h3><a id="user-content---" class="anchor" aria-label="永久链接：" href="#--"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<table>
<thead>
<tr>
<th><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">❗  </font></font><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">免责声明</font></font></strong></th>
</tr>
</thead>
<tbody>
<tr>
<td><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">该项目正在积极开发中</font></font></strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">。</font><font style="vertical-align: inherit;">预计发布后会有重大变化。</font><font style="vertical-align: inherit;">更新之前查看发布变更日志。</font></font></td>
</tr>
<tr>
<td><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">该项目主要是为了用作独立的 CLI 工具而构建的。</font></font><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">将 nucles 作为服务运行可能会带来安全风险。</font></font></strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">建议谨慎使用并采取额外的安全措施。</font></font></td>
</tr>
</tbody>
</table>
<div class="markdown-heading" dir="auto"><h1 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">安装核</font></font></h1><a id="user-content-install-nuclei" class="anchor" aria-label="永久链接：安装 Nuclei" href="#install-nuclei"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Nuclei 需要</font></font><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">go1.21</font></font></strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">才能成功安装。</font><font style="vertical-align: inherit;">运行以下命令来安装最新版本 -</font></font></p>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 tooltipped-no-delay d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="go install -v github.com/projectdiscovery/nuclei/v3/cmd/nuclei@latest" tabindex="0" role="button">
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon">
    <path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path>
</svg>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none">
    <path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0Z"></path>
</svg>
    </clipboard-copy>
  </div></div>
<details>
  <summary><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">酿造</font></font></summary>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>brew install nuclei</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 tooltipped-no-delay d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="brew install nuclei" tabindex="0" role="button">
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon">
    <path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path>
</svg>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none">
    <path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0Z"></path>
</svg>
    </clipboard-copy>
  </div></div>
</details>
<details>
  <summary><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">码头工人</font></font></summary>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>docker pull projectdiscovery/nuclei:latest</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 tooltipped-no-delay d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="docker pull projectdiscovery/nuclei:latest" tabindex="0" role="button">
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon">
    <path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path>
</svg>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none">
    <path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0Z"></path>
</svg>
    </clipboard-copy>
  </div></div>
</details>
<p dir="auto"><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">更多安装</font></font><a href="https://docs.projectdiscovery.io/tools/nuclei/install" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">方法可以在这里找到</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">。</font></font></strong></p>
<table>
<tbody><tr>
<td>  
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">细胞核模板</font></font></h3><a id="user-content-nuclei-templates" class="anchor" aria-label="永久链接：核模板" href="#nuclei-templates"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"></font><a href="https://github.com/projectdiscovery/nuclei/releases/tag/v2.5.2"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">自v2.5.2</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">版本以来，Nuclei 默认内置自动模板下载/更新支持</font><font style="vertical-align: inherit;">。</font></font><a href="https://github.com/projectdiscovery/nuclei-templates"><strong><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Nuclei-Templates</font></font></strong></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">项目提供了社区贡献的随时可用的模板列表，该列表不断更新。</font></font></p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">您仍然可以</font></font><code>update-templates</code><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">随时使用该标志来更新核模板；</font></font><a href="https://docs.projectdiscovery.io/templates/" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">您可以按照 Nuclei 的模板指南</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">为您的个人工作流程和需求编写自己的检查</font><font style="vertical-align: inherit;">。</font></font></p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">YAML DSL 参考语法可</font></font><a href="/projectdiscovery/nuclei/blob/dev/SYNTAX-REFERENCE.md"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">在此处</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">获取。</font></font></p>
</td>
</tr>
</tbody></table>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">用法</font></font></h3><a id="user-content-usage" class="anchor" aria-label="永久链接：用法" href="#usage"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>nuclei -h</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 tooltipped-no-delay d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="nuclei -h" tabindex="0" role="button">
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon">
    <path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path>
</svg>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none">
    <path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0Z"></path>
</svg>
    </clipboard-copy>
  </div></div>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">这将显示该工具的帮助。</font><font style="vertical-align: inherit;">这是它支持的所有交换机。</font></font></p>
<div class="highlight highlight-text-shell-session notranslate position-relative overflow-auto" dir="auto"><pre><span class="pl-c1">Nuclei is a fast, template based vulnerability scanner focusing</span>
<span class="pl-c1">on extensive configurability, massive extensibility and ease of use.</span>

<span class="pl-c1">Usage:</span>
<span class="pl-c1">  ./nuclei [flags]</span>

<span class="pl-c1">Flags:</span>
<span class="pl-c1">TARGET:</span>
<span class="pl-c1">   -u, -target string[]          target URLs/hosts to scan</span>
<span class="pl-c1">   -l, -list string              path to file containing a list of target URLs/hosts to scan (one per line)</span>
<span class="pl-c1">   -eh, -exclude-hosts string[]  hosts to exclude to scan from the input list (ip, cidr, hostname)</span>
<span class="pl-c1">   -resume string                resume scan using resume.cfg (clustering will be disabled)</span>
<span class="pl-c1">   -sa, -scan-all-ips            scan all the IP's associated with dns record</span>
<span class="pl-c1">   -iv, -ip-version string[]     IP version to scan of hostname (4,6) - (default 4)</span>

<span class="pl-c1">TARGET-FORMAT:</span>
<span class="pl-c1">   -im, -input-mode string        mode of input file (list, burp, jsonl, yaml, openapi, swagger) (default "list")</span>
<span class="pl-c1">   -ro, -required-only            use only required fields in input format when generating requests</span>
<span class="pl-c1">   -sfv, -skip-format-validation  skip format validation (like missing vars) when parsing input file</span>

<span class="pl-c1">TEMPLATES:</span>
<span class="pl-c1">   -nt, -new-templates                    run only new templates added in latest nuclei-templates release</span>
<span class="pl-c1">   -ntv, -new-templates-version string[]  run new templates added in specific version</span>
<span class="pl-c1">   -as, -automatic-scan                   automatic web scan using wappalyzer technology detection to tags mapping</span>
<span class="pl-c1">   -t, -templates string[]                list of template or template directory to run (comma-separated, file)</span>
<span class="pl-c1">   -turl, -template-url string[]          template url or list containing template urls to run (comma-separated, file)</span>
<span class="pl-c1">   -w, -workflows string[]                list of workflow or workflow directory to run (comma-separated, file)</span>
<span class="pl-c1">   -wurl, -workflow-url string[]          workflow url or list containing workflow urls to run (comma-separated, file)</span>
<span class="pl-c1">   -validate                              validate the passed templates to nuclei</span>
<span class="pl-c1">   -nss, -no-strict-syntax                disable strict syntax check on templates</span>
<span class="pl-c1">   -td, -template-display                 displays the templates content</span>
<span class="pl-c1">   -tl                                    list all available templates</span>
<span class="pl-c1">   -sign                                  signs the templates with the private key defined in NUCLEI_SIGNATURE_PRIVATE_KEY env variable</span>
<span class="pl-c1">   -code                                  enable loading code protocol-based templates</span>
<span class="pl-c1">   -dut, -disable-unsigned-templates      disable running unsigned templates or templates with mismatched signature</span>

<span class="pl-c1">FILTERING:</span>
<span class="pl-c1">   -a, -author string[]               templates to run based on authors (comma-separated, file)</span>
<span class="pl-c1">   -tags string[]                     templates to run based on tags (comma-separated, file)</span>
<span class="pl-c1">   -etags, -exclude-tags string[]     templates to exclude based on tags (comma-separated, file)</span>
<span class="pl-c1">   -itags, -include-tags string[]     tags to be executed even if they are excluded either by default or configuration</span>
<span class="pl-c1">   -id, -template-id string[]         templates to run based on template ids (comma-separated, file, allow-wildcard)</span>
<span class="pl-c1">   -eid, -exclude-id string[]         templates to exclude based on template ids (comma-separated, file)</span>
<span class="pl-c1">   -it, -include-templates string[]   path to template file or directory to be executed even if they are excluded either by default or configuration</span>
<span class="pl-c1">   -et, -exclude-templates string[]   path to template file or directory to exclude (comma-separated, file)</span>
<span class="pl-c1">   -em, -exclude-matchers string[]    template matchers to exclude in result</span>
<span class="pl-c1">   -s, -severity value[]              templates to run based on severity. Possible values: info, low, medium, high, critical, unknown</span>
<span class="pl-c1">   -es, -exclude-severity value[]     templates to exclude based on severity. Possible values: info, low, medium, high, critical, unknown</span>
<span class="pl-c1">   -pt, -type value[]                 templates to run based on protocol type. Possible values: dns, file, http, headless, tcp, workflow, ssl, websocket, whois, code, javascript</span>
<span class="pl-c1">   -ept, -exclude-type value[]        templates to exclude based on protocol type. Possible values: dns, file, http, headless, tcp, workflow, ssl, websocket, whois, code, javascript</span>
<span class="pl-c1">   -tc, -template-condition string[]  templates to run based on expression condition</span>

<span class="pl-c1">OUTPUT:</span>
<span class="pl-c1">   -o, -output string            output file to write found issues/vulnerabilities</span>
<span class="pl-c1">   -sresp, -store-resp           store all request/response passed through nuclei to output directory</span>
<span class="pl-c1">   -srd, -store-resp-dir string  store all request/response passed through nuclei to custom directory (default "output")</span>
<span class="pl-c1">   -silent                       display findings only</span>
<span class="pl-c1">   -nc, -no-color                disable output content coloring (ANSI escape codes)</span>
<span class="pl-c1">   -j, -jsonl                    write output in JSONL(ines) format</span>
<span class="pl-c1">   -irr, -include-rr -omit-raw   include request/response pairs in the JSON, JSONL, and Markdown outputs (for findings only) [DEPRECATED use -omit-raw] (default true)</span>
<span class="pl-c1">   -or, -omit-raw                omit request/response pairs in the JSON, JSONL, and Markdown outputs (for findings only)</span>
<span class="pl-c1">   -ot, -omit-template           omit encoded template in the JSON, JSONL output</span>
<span class="pl-c1">   -nm, -no-meta                 disable printing result metadata in cli output</span>
<span class="pl-c1">   -ts, -timestamp               enables printing timestamp in cli output</span>
<span class="pl-c1">   -rdb, -report-db string       nuclei reporting database (always use this to persist report data)</span>
<span class="pl-c1">   -ms, -matcher-status          display match failure status</span>
<span class="pl-c1">   -me, -markdown-export string  directory to export results in markdown format</span>
<span class="pl-c1">   -se, -sarif-export string     file to export results in SARIF format</span>
<span class="pl-c1">   -je, -json-export string      file to export results in JSON format</span>
<span class="pl-c1">   -jle, -jsonl-export string    file to export results in JSONL(ine) format</span>

<span class="pl-c1">CONFIGURATIONS:</span>
<span class="pl-c1">   -config string                        path to the nuclei configuration file</span>
<span class="pl-c1">   -fr, -follow-redirects                enable following redirects for http templates</span>
<span class="pl-c1">   -fhr, -follow-host-redirects          follow redirects on the same host</span>
<span class="pl-c1">   -mr, -max-redirects int               max number of redirects to follow for http templates (default 10)</span>
<span class="pl-c1">   -dr, -disable-redirects               disable redirects for http templates</span>
<span class="pl-c1">   -rc, -report-config string            nuclei reporting module configuration file</span>
<span class="pl-c1">   -H, -header string[]                  custom header/cookie to include in all http request in header:value format (cli, file)</span>
<span class="pl-c1">   -V, -var value                        custom vars in key=value format</span>
<span class="pl-c1">   -r, -resolvers string                 file containing resolver list for nuclei</span>
<span class="pl-c1">   -sr, -system-resolvers                use system DNS resolving as error fallback</span>
<span class="pl-c1">   -dc, -disable-clustering              disable clustering of requests</span>
<span class="pl-c1">   -passive                              enable passive HTTP response processing mode</span>
<span class="pl-c1">   -fh2, -force-http2                    force http2 connection on requests</span>
<span class="pl-c1">   -ev, -env-vars                        enable environment variables to be used in template</span>
<span class="pl-c1">   -cc, -client-cert string              client certificate file (PEM-encoded) used for authenticating against scanned hosts</span>
<span class="pl-c1">   -ck, -client-key string               client key file (PEM-encoded) used for authenticating against scanned hosts</span>
<span class="pl-c1">   -ca, -client-ca string                client certificate authority file (PEM-encoded) used for authenticating against scanned hosts</span>
<span class="pl-c1">   -sml, -show-match-line                show match lines for file templates, works with extractors only</span>
<span class="pl-c1">   -ztls                                 use ztls library with autofallback to standard one for tls13 [Deprecated] autofallback to ztls is enabled by default</span>
<span class="pl-c1">   -sni string                           tls sni hostname to use (default: input domain name)</span>
<span class="pl-c1">   -dt, -dialer-timeout value            timeout for network requests.</span>
<span class="pl-c1">   -dka, -dialer-keep-alive value        keep-alive duration for network requests.</span>
<span class="pl-c1">   -lfa, -allow-local-file-access        allows file (payload) access anywhere on the system</span>
<span class="pl-c1">   -lna, -restrict-local-network-access  blocks connections to the local / private network</span>
<span class="pl-c1">   -i, -interface string                 network interface to use for network scan</span>
<span class="pl-c1">   -at, -attack-type string              type of payload combinations to perform (batteringram,pitchfork,clusterbomb)</span>
<span class="pl-c1">   -sip, -source-ip string               source ip address to use for network scan</span>
<span class="pl-c1">   -rsr, -response-size-read int         max response size to read in bytes (default 10485760)</span>
<span class="pl-c1">   -rss, -response-size-save int         max response size to read in bytes (default 1048576)</span>
<span class="pl-c1">   -reset                                reset removes all nuclei configuration and data files (including nuclei-templates)</span>
<span class="pl-c1">   -tlsi, -tls-impersonate               enable experimental client hello (ja3) tls randomization</span>

<span class="pl-c1">INTERACTSH:</span>
<span class="pl-c1">   -iserver, -interactsh-server string  interactsh server url for self-hosted instance (default: oast.pro,oast.live,oast.site,oast.online,oast.fun,oast.me)</span>
<span class="pl-c1">   -itoken, -interactsh-token string    authentication token for self-hosted interactsh server</span>
<span class="pl-c1">   -interactions-cache-size int         number of requests to keep in the interactions cache (default 5000)</span>
<span class="pl-c1">   -interactions-eviction int           number of seconds to wait before evicting requests from cache (default 60)</span>
<span class="pl-c1">   -interactions-poll-duration int      number of seconds to wait before each interaction poll request (default 5)</span>
<span class="pl-c1">   -interactions-cooldown-period int    extra time for interaction polling before exiting (default 5)</span>
<span class="pl-c1">   -ni, -no-interactsh                  disable interactsh server for OAST testing, exclude OAST based templates</span>

<span class="pl-c1">FUZZING:</span>
<span class="pl-c1">   -ft, -fuzzing-type string  overrides fuzzing type set in template (replace, prefix, postfix, infix)</span>
<span class="pl-c1">   -fm, -fuzzing-mode string  overrides fuzzing mode set in template (multiple, single)</span>
<span class="pl-c1">   -fuzz                      enable loading fuzzing templates</span>

<span class="pl-c1">UNCOVER:</span>
<span class="pl-c1">   -uc, -uncover                  enable uncover engine</span>
<span class="pl-c1">   -uq, -uncover-query string[]   uncover search query</span>
<span class="pl-c1">   -ue, -uncover-engine string[]  uncover search engine (shodan,censys,fofa,shodan-idb,quake,hunter,zoomeye,netlas,criminalip,publicwww,hunterhow) (default shodan)</span>
<span class="pl-c1">   -uf, -uncover-field string     uncover fields to return (ip,port,host) (default "ip:port")</span>
<span class="pl-c1">   -ul, -uncover-limit int        uncover results to return (default 100)</span>
<span class="pl-c1">   -ur, -uncover-ratelimit int    override ratelimit of engines with unknown ratelimit (default 60 req/min) (default 60)</span>

<span class="pl-c1">RATE-LIMIT:</span>
<span class="pl-c1">   -rl, -rate-limit int               maximum number of requests to send per second (default 150)</span>
<span class="pl-c1">   -rlm, -rate-limit-minute int       maximum number of requests to send per minute</span>
<span class="pl-c1">   -bs, -bulk-size int                maximum number of hosts to be analyzed in parallel per template (default 25)</span>
<span class="pl-c1">   -c, -concurrency int               maximum number of templates to be executed in parallel (default 25)</span>
<span class="pl-c1">   -hbs, -headless-bulk-size int      maximum number of headless hosts to be analyzed in parallel per template (default 10)</span>
<span class="pl-c1">   -headc, -headless-concurrency int  maximum number of headless templates to be executed in parallel (default 10)</span>
<span class="pl-c1">   -jsc, -js-concurrency int          maximum number of javascript runtimes to be executed in parallel (default 120)</span>
<span class="pl-c1">   -pc, -payload-concurrency int      max payload concurrency for each template (default 25)</span>

<span class="pl-c1">OPTIMIZATIONS:</span>
<span class="pl-c1">   -timeout int                     time to wait in seconds before timeout (default 10)</span>
<span class="pl-c1">   -retries int                     number of times to retry a failed request (default 1)</span>
<span class="pl-c1">   -ldp, -leave-default-ports       leave default HTTP/HTTPS ports (eg. host:80,host:443)</span>
<span class="pl-c1">   -mhe, -max-host-error int        max errors for a host before skipping from scan (default 30)</span>
<span class="pl-c1">   -te, -track-error string[]       adds given error to max-host-error watchlist (standard, file)</span>
<span class="pl-c1">   -nmhe, -no-mhe                   disable skipping host from scan based on errors</span>
<span class="pl-c1">   -project                         use a project folder to avoid sending same request multiple times</span>
<span class="pl-c1">   -project-path string             set a specific project path (default "/tmp")</span>
<span class="pl-c1">   -spm, -stop-at-first-match       stop processing HTTP requests after the first match (may break template/workflow logic)</span>
<span class="pl-c1">   -stream                          stream mode - start elaborating without sorting the input</span>
<span class="pl-c1">   -ss, -scan-strategy value        strategy to use while scanning(auto/host-spray/template-spray) (default auto)</span>
<span class="pl-c1">   -irt, -input-read-timeout value  timeout on input read (default 3m0s)</span>
<span class="pl-c1">   -nh, -no-httpx                   disable httpx probing for non-url input</span>
<span class="pl-c1">   -no-stdin                        disable stdin processing</span>

<span class="pl-c1">HEADLESS:</span>
<span class="pl-c1">   -headless                        enable templates that require headless browser support (root user on Linux will disable sandbox)</span>
<span class="pl-c1">   -page-timeout int                seconds to wait for each page in headless mode (default 20)</span>
<span class="pl-c1">   -sb, -show-browser               show the browser on the screen when running templates with headless mode</span>
<span class="pl-c1">   -ho, -headless-options string[]  start headless chrome with additional options</span>
<span class="pl-c1">   -sc, -system-chrome              use local installed Chrome browser instead of nuclei installed</span>
<span class="pl-c1">   -lha, -list-headless-action      list available headless actions</span>

<span class="pl-c1">DEBUG:</span>
<span class="pl-c1">   -debug                    show all requests and responses</span>
<span class="pl-c1">   -dreq, -debug-req         show all sent requests</span>
<span class="pl-c1">   -dresp, -debug-resp       show all received responses</span>
<span class="pl-c1">   -p, -proxy string[]       list of http/socks5 proxy to use (comma separated or file input)</span>
<span class="pl-c1">   -pi, -proxy-internal      proxy all internal requests</span>
<span class="pl-c1">   -ldf, -list-dsl-function  list all supported DSL function signatures</span>
<span class="pl-c1">   -tlog, -trace-log string  file to write sent requests trace log</span>
<span class="pl-c1">   -elog, -error-log string  file to write sent requests error log</span>
<span class="pl-c1">   -version                  show nuclei version</span>
<span class="pl-c1">   -hm, -hang-monitor        enable nuclei hang monitoring</span>
<span class="pl-c1">   -v, -verbose              show verbose output</span>
<span class="pl-c1">   -profile-mem string       optional nuclei memory profile dump file</span>
<span class="pl-c1">   -vv                       display templates loaded for scan</span>
<span class="pl-c1">   -svd, -show-var-dump      show variables dump for debugging</span>
<span class="pl-c1">   -ep, -enable-pprof        enable pprof debugging server</span>
<span class="pl-c1">   -tv, -templates-version   shows the version of the installed nuclei-templates</span>
<span class="pl-c1">   -hc, -health-check        run diagnostic check up</span>

<span class="pl-c1">UPDATE:</span>
<span class="pl-c1">   -up, -update                      update nuclei engine to the latest released version</span>
<span class="pl-c1">   -ut, -update-templates            update nuclei-templates to latest released version</span>
<span class="pl-c1">   -ud, -update-template-dir string  custom directory to install / update nuclei-templates</span>
<span class="pl-c1">   -duc, -disable-update-check       disable automatic nuclei/templates update check</span>

<span class="pl-c1">STATISTICS:</span>
<span class="pl-c1">   -stats                    display statistics about the running scan</span>
<span class="pl-c1">   -sj, -stats-json          display statistics in JSONL(ines) format</span>
<span class="pl-c1">   -si, -stats-interval int  number of seconds to wait between showing a statistics update (default 5)</span>
<span class="pl-c1">   -mp, -metrics-port int    port to expose nuclei metrics on (default 9092)</span>

<span class="pl-c1">CLOUD:</span>
<span class="pl-c1">   -auth                  configure projectdiscovery cloud (pdcp) api key</span>
<span class="pl-c1">   -cup, -cloud-upload    upload scan results to pdcp dashboard</span>
<span class="pl-c1">   -sid, -scan-id string  upload scan results to given scan id</span>

<span class="pl-c1">AUTHENTICATION:</span>
<span class="pl-c1">   -sf, -secret-file string[]  path to config file containing secrets for nuclei authenticated scan</span>
<span class="pl-c1">   -ps, -prefetch-secrets      prefetch secrets from the secrets file</span>


<span class="pl-c1">EXAMPLES:</span>
<span class="pl-c1">Run nuclei on single host:</span>
<span class="pl-c1">   $ nuclei -target example.com</span>

<span class="pl-c1">Run nuclei with specific template directories:</span>
<span class="pl-c1">   $ nuclei -target example.com -t http/cves/ -t ssl</span>

<span class="pl-c1">Run nuclei against a list of hosts:</span>
<span class="pl-c1">   $ nuclei -list hosts.txt</span>

<span class="pl-c1">Run nuclei with a JSON output:</span>
<span class="pl-c1">   $ nuclei -target example.com -json-export output.json</span>

<span class="pl-c1">Run nuclei with sorted Markdown outputs (with environment variables):</span>
<span class="pl-c1">   $ MARKDOWN_EXPORT_SORT_MODE=template nuclei -target example.com -markdown-export nuclei_report/</span>

<span class="pl-c1">Additional documentation is available at: https://docs.nuclei.sh/getting-started/running</span></pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 tooltipped-no-delay d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="Nuclei is a fast, template based vulnerability scanner focusing
on extensive configurability, massive extensibility and ease of use.

Usage:
  ./nuclei [flags]

Flags:
TARGET:
   -u, -target string[]          target URLs/hosts to scan
   -l, -list string              path to file containing a list of target URLs/hosts to scan (one per line)
   -eh, -exclude-hosts string[]  hosts to exclude to scan from the input list (ip, cidr, hostname)
   -resume string                resume scan using resume.cfg (clustering will be disabled)
   -sa, -scan-all-ips            scan all the IP's associated with dns record
   -iv, -ip-version string[]     IP version to scan of hostname (4,6) - (default 4)

TARGET-FORMAT:
   -im, -input-mode string        mode of input file (list, burp, jsonl, yaml, openapi, swagger) (default &quot;list&quot;)
   -ro, -required-only            use only required fields in input format when generating requests
   -sfv, -skip-format-validation  skip format validation (like missing vars) when parsing input file

TEMPLATES:
   -nt, -new-templates                    run only new templates added in latest nuclei-templates release
   -ntv, -new-templates-version string[]  run new templates added in specific version
   -as, -automatic-scan                   automatic web scan using wappalyzer technology detection to tags mapping
   -t, -templates string[]                list of template or template directory to run (comma-separated, file)
   -turl, -template-url string[]          template url or list containing template urls to run (comma-separated, file)
   -w, -workflows string[]                list of workflow or workflow directory to run (comma-separated, file)
   -wurl, -workflow-url string[]          workflow url or list containing workflow urls to run (comma-separated, file)
   -validate                              validate the passed templates to nuclei
   -nss, -no-strict-syntax                disable strict syntax check on templates
   -td, -template-display                 displays the templates content
   -tl                                    list all available templates
   -sign                                  signs the templates with the private key defined in NUCLEI_SIGNATURE_PRIVATE_KEY env variable
   -code                                  enable loading code protocol-based templates
   -dut, -disable-unsigned-templates      disable running unsigned templates or templates with mismatched signature

FILTERING:
   -a, -author string[]               templates to run based on authors (comma-separated, file)
   -tags string[]                     templates to run based on tags (comma-separated, file)
   -etags, -exclude-tags string[]     templates to exclude based on tags (comma-separated, file)
   -itags, -include-tags string[]     tags to be executed even if they are excluded either by default or configuration
   -id, -template-id string[]         templates to run based on template ids (comma-separated, file, allow-wildcard)
   -eid, -exclude-id string[]         templates to exclude based on template ids (comma-separated, file)
   -it, -include-templates string[]   path to template file or directory to be executed even if they are excluded either by default or configuration
   -et, -exclude-templates string[]   path to template file or directory to exclude (comma-separated, file)
   -em, -exclude-matchers string[]    template matchers to exclude in result
   -s, -severity value[]              templates to run based on severity. Possible values: info, low, medium, high, critical, unknown
   -es, -exclude-severity value[]     templates to exclude based on severity. Possible values: info, low, medium, high, critical, unknown
   -pt, -type value[]                 templates to run based on protocol type. Possible values: dns, file, http, headless, tcp, workflow, ssl, websocket, whois, code, javascript
   -ept, -exclude-type value[]        templates to exclude based on protocol type. Possible values: dns, file, http, headless, tcp, workflow, ssl, websocket, whois, code, javascript
   -tc, -template-condition string[]  templates to run based on expression condition

OUTPUT:
   -o, -output string            output file to write found issues/vulnerabilities
   -sresp, -store-resp           store all request/response passed through nuclei to output directory
   -srd, -store-resp-dir string  store all request/response passed through nuclei to custom directory (default &quot;output&quot;)
   -silent                       display findings only
   -nc, -no-color                disable output content coloring (ANSI escape codes)
   -j, -jsonl                    write output in JSONL(ines) format
   -irr, -include-rr -omit-raw   include request/response pairs in the JSON, JSONL, and Markdown outputs (for findings only) [DEPRECATED use -omit-raw] (default true)
   -or, -omit-raw                omit request/response pairs in the JSON, JSONL, and Markdown outputs (for findings only)
   -ot, -omit-template           omit encoded template in the JSON, JSONL output
   -nm, -no-meta                 disable printing result metadata in cli output
   -ts, -timestamp               enables printing timestamp in cli output
   -rdb, -report-db string       nuclei reporting database (always use this to persist report data)
   -ms, -matcher-status          display match failure status
   -me, -markdown-export string  directory to export results in markdown format
   -se, -sarif-export string     file to export results in SARIF format
   -je, -json-export string      file to export results in JSON format
   -jle, -jsonl-export string    file to export results in JSONL(ine) format

CONFIGURATIONS:
   -config string                        path to the nuclei configuration file
   -fr, -follow-redirects                enable following redirects for http templates
   -fhr, -follow-host-redirects          follow redirects on the same host
   -mr, -max-redirects int               max number of redirects to follow for http templates (default 10)
   -dr, -disable-redirects               disable redirects for http templates
   -rc, -report-config string            nuclei reporting module configuration file
   -H, -header string[]                  custom header/cookie to include in all http request in header:value format (cli, file)
   -V, -var value                        custom vars in key=value format
   -r, -resolvers string                 file containing resolver list for nuclei
   -sr, -system-resolvers                use system DNS resolving as error fallback
   -dc, -disable-clustering              disable clustering of requests
   -passive                              enable passive HTTP response processing mode
   -fh2, -force-http2                    force http2 connection on requests
   -ev, -env-vars                        enable environment variables to be used in template
   -cc, -client-cert string              client certificate file (PEM-encoded) used for authenticating against scanned hosts
   -ck, -client-key string               client key file (PEM-encoded) used for authenticating against scanned hosts
   -ca, -client-ca string                client certificate authority file (PEM-encoded) used for authenticating against scanned hosts
   -sml, -show-match-line                show match lines for file templates, works with extractors only
   -ztls                                 use ztls library with autofallback to standard one for tls13 [Deprecated] autofallback to ztls is enabled by default
   -sni string                           tls sni hostname to use (default: input domain name)
   -dt, -dialer-timeout value            timeout for network requests.
   -dka, -dialer-keep-alive value        keep-alive duration for network requests.
   -lfa, -allow-local-file-access        allows file (payload) access anywhere on the system
   -lna, -restrict-local-network-access  blocks connections to the local / private network
   -i, -interface string                 network interface to use for network scan
   -at, -attack-type string              type of payload combinations to perform (batteringram,pitchfork,clusterbomb)
   -sip, -source-ip string               source ip address to use for network scan
   -rsr, -response-size-read int         max response size to read in bytes (default 10485760)
   -rss, -response-size-save int         max response size to read in bytes (default 1048576)
   -reset                                reset removes all nuclei configuration and data files (including nuclei-templates)
   -tlsi, -tls-impersonate               enable experimental client hello (ja3) tls randomization

INTERACTSH:
   -iserver, -interactsh-server string  interactsh server url for self-hosted instance (default: oast.pro,oast.live,oast.site,oast.online,oast.fun,oast.me)
   -itoken, -interactsh-token string    authentication token for self-hosted interactsh server
   -interactions-cache-size int         number of requests to keep in the interactions cache (default 5000)
   -interactions-eviction int           number of seconds to wait before evicting requests from cache (default 60)
   -interactions-poll-duration int      number of seconds to wait before each interaction poll request (default 5)
   -interactions-cooldown-period int    extra time for interaction polling before exiting (default 5)
   -ni, -no-interactsh                  disable interactsh server for OAST testing, exclude OAST based templates

FUZZING:
   -ft, -fuzzing-type string  overrides fuzzing type set in template (replace, prefix, postfix, infix)
   -fm, -fuzzing-mode string  overrides fuzzing mode set in template (multiple, single)
   -fuzz                      enable loading fuzzing templates

UNCOVER:
   -uc, -uncover                  enable uncover engine
   -uq, -uncover-query string[]   uncover search query
   -ue, -uncover-engine string[]  uncover search engine (shodan,censys,fofa,shodan-idb,quake,hunter,zoomeye,netlas,criminalip,publicwww,hunterhow) (default shodan)
   -uf, -uncover-field string     uncover fields to return (ip,port,host) (default &quot;ip:port&quot;)
   -ul, -uncover-limit int        uncover results to return (default 100)
   -ur, -uncover-ratelimit int    override ratelimit of engines with unknown ratelimit (default 60 req/min) (default 60)

RATE-LIMIT:
   -rl, -rate-limit int               maximum number of requests to send per second (default 150)
   -rlm, -rate-limit-minute int       maximum number of requests to send per minute
   -bs, -bulk-size int                maximum number of hosts to be analyzed in parallel per template (default 25)
   -c, -concurrency int               maximum number of templates to be executed in parallel (default 25)
   -hbs, -headless-bulk-size int      maximum number of headless hosts to be analyzed in parallel per template (default 10)
   -headc, -headless-concurrency int  maximum number of headless templates to be executed in parallel (default 10)
   -jsc, -js-concurrency int          maximum number of javascript runtimes to be executed in parallel (default 120)
   -pc, -payload-concurrency int      max payload concurrency for each template (default 25)

OPTIMIZATIONS:
   -timeout int                     time to wait in seconds before timeout (default 10)
   -retries int                     number of times to retry a failed request (default 1)
   -ldp, -leave-default-ports       leave default HTTP/HTTPS ports (eg. host:80,host:443)
   -mhe, -max-host-error int        max errors for a host before skipping from scan (default 30)
   -te, -track-error string[]       adds given error to max-host-error watchlist (standard, file)
   -nmhe, -no-mhe                   disable skipping host from scan based on errors
   -project                         use a project folder to avoid sending same request multiple times
   -project-path string             set a specific project path (default &quot;/tmp&quot;)
   -spm, -stop-at-first-match       stop processing HTTP requests after the first match (may break template/workflow logic)
   -stream                          stream mode - start elaborating without sorting the input
   -ss, -scan-strategy value        strategy to use while scanning(auto/host-spray/template-spray) (default auto)
   -irt, -input-read-timeout value  timeout on input read (default 3m0s)
   -nh, -no-httpx                   disable httpx probing for non-url input
   -no-stdin                        disable stdin processing

HEADLESS:
   -headless                        enable templates that require headless browser support (root user on Linux will disable sandbox)
   -page-timeout int                seconds to wait for each page in headless mode (default 20)
   -sb, -show-browser               show the browser on the screen when running templates with headless mode
   -ho, -headless-options string[]  start headless chrome with additional options
   -sc, -system-chrome              use local installed Chrome browser instead of nuclei installed
   -lha, -list-headless-action      list available headless actions

DEBUG:
   -debug                    show all requests and responses
   -dreq, -debug-req         show all sent requests
   -dresp, -debug-resp       show all received responses
   -p, -proxy string[]       list of http/socks5 proxy to use (comma separated or file input)
   -pi, -proxy-internal      proxy all internal requests
   -ldf, -list-dsl-function  list all supported DSL function signatures
   -tlog, -trace-log string  file to write sent requests trace log
   -elog, -error-log string  file to write sent requests error log
   -version                  show nuclei version
   -hm, -hang-monitor        enable nuclei hang monitoring
   -v, -verbose              show verbose output
   -profile-mem string       optional nuclei memory profile dump file
   -vv                       display templates loaded for scan
   -svd, -show-var-dump      show variables dump for debugging
   -ep, -enable-pprof        enable pprof debugging server
   -tv, -templates-version   shows the version of the installed nuclei-templates
   -hc, -health-check        run diagnostic check up

UPDATE:
   -up, -update                      update nuclei engine to the latest released version
   -ut, -update-templates            update nuclei-templates to latest released version
   -ud, -update-template-dir string  custom directory to install / update nuclei-templates
   -duc, -disable-update-check       disable automatic nuclei/templates update check

STATISTICS:
   -stats                    display statistics about the running scan
   -sj, -stats-json          display statistics in JSONL(ines) format
   -si, -stats-interval int  number of seconds to wait between showing a statistics update (default 5)
   -mp, -metrics-port int    port to expose nuclei metrics on (default 9092)

CLOUD:
   -auth                  configure projectdiscovery cloud (pdcp) api key
   -cup, -cloud-upload    upload scan results to pdcp dashboard
   -sid, -scan-id string  upload scan results to given scan id

AUTHENTICATION:
   -sf, -secret-file string[]  path to config file containing secrets for nuclei authenticated scan
   -ps, -prefetch-secrets      prefetch secrets from the secrets file


EXAMPLES:
Run nuclei on single host:
   $ nuclei -target example.com

Run nuclei with specific template directories:
   $ nuclei -target example.com -t http/cves/ -t ssl

Run nuclei against a list of hosts:
   $ nuclei -list hosts.txt

Run nuclei with a JSON output:
   $ nuclei -target example.com -json-export output.json

Run nuclei with sorted Markdown outputs (with environment variables):
   $ MARKDOWN_EXPORT_SORT_MODE=template nuclei -target example.com -markdown-export nuclei_report/

Additional documentation is available at: https://docs.nuclei.sh/getting-started/running" tabindex="0" role="button">
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-copy js-clipboard-copy-icon">
    <path d="M0 6.75C0 5.784.784 5 1.75 5h1.5a.75.75 0 0 1 0 1.5h-1.5a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-1.5a.75.75 0 0 1 1.5 0v1.5A1.75 1.75 0 0 1 9.25 16h-7.5A1.75 1.75 0 0 1 0 14.25Z"></path><path d="M5 1.75C5 .784 5.784 0 6.75 0h7.5C15.216 0 16 .784 16 1.75v7.5A1.75 1.75 0 0 1 14.25 11h-7.5A1.75 1.75 0 0 1 5 9.25Zm1.75-.25a.25.25 0 0 0-.25.25v7.5c0 .138.112.25.25.25h7.5a.25.25 0 0 0 .25-.25v-7.5a.25.25 0 0 0-.25-.25Z"></path>
</svg>
      <svg aria-hidden="true" height="16" viewBox="0 0 16 16" version="1.1" width="16" data-view-component="true" class="octicon octicon-check js-clipboard-check-icon color-fg-success d-none">
    <path d="M13.78 4.22a.75.75 0 0 1 0 1.06l-7.25 7.25a.75.75 0 0 1-1.06 0L2.22 9.28a.751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018L6 10.94l6.72-6.72a.75.75 0 0 1 1.06 0Z"></path>
</svg>
    </clipboard-copy>
  </div></div>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">运行核</font></font></h3><a id="user-content-running-nuclei" class="anchor" aria-label="永久链接：运行核" href="#running-nuclei"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">有关运行 Nuclei 的详细信息，</font><font style="vertical-align: inherit;">请参阅</font></font><a href="https://docs.projectdiscovery.io/tools/nuclei/running" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">https://docs.projectdiscovery.io/tools/nuclei/running</font></font></a><font style="vertical-align: inherit;"></font></p>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">使用 Go 代码中的 Nuclei</font></font></h3><a id="user-content-using-nuclei-from-go-code" class="anchor" aria-label="永久链接：使用 Go 代码中的 Nuclei" href="#using-nuclei-from-go-code"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"><a href="https://pkg.go.dev/github.com/projectdiscovery/nuclei/v3/lib#section-readme" rel="nofollow"><font style="vertical-align: inherit;">godoc</font></a><font style="vertical-align: inherit;">提供了使用 Nuclei 作为库/SDK 的完整指南</font></font><a href="https://pkg.go.dev/github.com/projectdiscovery/nuclei/v3/lib#section-readme" rel="nofollow"><font style="vertical-align: inherit;"></font></a></p>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">资源</font></font></h3><a id="user-content-resources" class="anchor" aria-label="永久链接：资源" href="#resources"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"></font><a href="https://docs.projectdiscovery.io/tools/nuclei/" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">您可以在https://docs.projectdiscovery.io/tools/nuclei/</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">访问 Nuclei 的主要文档，并通过</font><a href="https://cloud.projectdiscovery.io" rel="nofollow"><font style="vertical-align: inherit;">ProjectDiscovery Cloud Platform</font></a><font style="vertical-align: inherit;">了解有关云中 Nuclei 的更多信息</font></font><a href="https://cloud.projectdiscovery.io" rel="nofollow"><font style="vertical-align: inherit;"></font></a></p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">有关 Nuclei 的更多资源和视频，</font><font style="vertical-align: inherit;">请参阅</font></font><a href="https://docs.projectdiscovery.io/tools/nuclei/resources" rel="nofollow"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">https://docs.projectdiscovery.io/tools/nuclei/resources ！</font></font></a><font style="vertical-align: inherit;"></font></p>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">制作人员</font></font></h3><a id="user-content-credits" class="anchor" aria-label="永久链接：学分" href="#credits"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">感谢所有出色的</font></font><a href="https://github.com/projectdiscovery/nuclei/graphs/contributors"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">社区贡献者发送 PR</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">并保持该项目的更新。</font><font style="vertical-align: inherit;">❤️</font></font></p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">如果您有想法或某种改进，欢迎您贡献并参与该项目，请随时发送您的 PR。</font></font></p>
<p align="center" dir="auto">
<a href="https://github.com/projectdiscovery/nuclei/graphs/contributors">
  <img src="https://camo.githubusercontent.com/77e00d39c32b1d0019049081bac2f8b792c17ddae4913868fa900362b6e8facb/68747470733a2f2f636f6e747269622e726f636b732f696d6167653f7265706f3d70726f6a656374646973636f766572792f6e75636c6569266d61783d353030" data-canonical-src="https://contrib.rocks/image?repo=projectdiscovery/nuclei&amp;max=500" style="max-width: 100%;">
</a>
</p>
<p dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">另请查看以下可能适合您的工作流程的类似开源项目：</font></font></p>
<p dir="auto"><a href="https://github.com/ffuf/ffuf"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">FFuF</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/ameenmaali/qsfuzz"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Qsfuzz</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/proabiral/inception"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Inception</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/hannob/snallygaster"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Snallygaster</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/Static-Flow/gofingerprint"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Gofingerprint</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/1N3/Sn1per/tree/master/templates"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Sn1per</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/google/tsunami-security-scanner"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Google 海啸</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/jaeles-project/jaeles"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">Jaeles</font></font></a><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">、</font></font><a href="https://github.com/michelin/ChopChop"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">ChopChop</font></font></a></p>
<div class="markdown-heading" dir="auto"><h3 tabindex="-1" class="heading-element" dir="auto"><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">执照</font></font></h3><a id="user-content-license" class="anchor" aria-label="永久链接：许可证" href="#license"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
<p dir="auto"><font style="vertical-align: inherit;"><a href="https://github.com/projectdiscovery/nuclei/blob/main/LICENSE.md"><font style="vertical-align: inherit;">Nuclei 根据MIT 许可证</font></a><font style="vertical-align: inherit;">分发</font></font><a href="https://github.com/projectdiscovery/nuclei/blob/main/LICENSE.md"><font style="vertical-align: inherit;"></font></a></p>
<div class="markdown-heading" dir="auto"><h1 align="left" tabindex="-1" class="heading-element" dir="auto">
  <a href="https://discord.gg/projectdiscovery" rel="nofollow"><img src="/projectdiscovery/nuclei/raw/dev/static/Join-Discord.png" width="380" alt="加入不和谐" style="max-width: 100%;"></a> <a href="https://docs.projectdiscovery.io" rel="nofollow"><img src="/projectdiscovery/nuclei/raw/dev/static/check-nuclei-documentation.png" width="380" alt="检查核文档" style="max-width: 100%;"></a>
</h1><a id="user-content----" class="anchor" aria-label="永久链接：" href="#---"><svg class="octicon octicon-link" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true"><path d="m7.775 3.275 1.25-1.25a3.5 3.5 0 1 1 4.95 4.95l-2.5 2.5a3.5 3.5 0 0 1-4.95 0 .751.751 0 0 1 .018-1.042.751.751 0 0 1 1.042-.018 1.998 1.998 0 0 0 2.83 0l2.5-2.5a2.002 2.002 0 0 0-2.83-2.83l-1.25 1.25a.751.751 0 0 1-1.042-.018.751.751 0 0 1-.018-1.042Zm-4.69 9.64a1.998 1.998 0 0 0 2.83 0l1.25-1.25a.751.751 0 0 1 1.042.018.751.751 0 0 1 .018 1.042l-1.25 1.25a3.5 3.5 0 1 1-4.95-4.95l2.5-2.5a3.5 3.5 0 0 1 4.95 0 .751.751 0 0 1-.018 1.042.751.751 0 0 1-1.042.018 1.998 1.998 0 0 0-2.83 0l-2.5 2.5a1.998 1.998 0 0 0 0 2.83Z"></path></svg></a></div>
</article></div>
