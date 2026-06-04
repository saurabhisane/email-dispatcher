# email-dispatcher

<div align="center">

<!-- README UI -->
<table width="100%" style="border-collapse:collapse;border:1px solid #2d333b;border-radius:12px;overflow:hidden;background:#0d1117;font-family:'Courier New',monospace;">
<tr style="background:#161b22;border-bottom:1px solid #2d333b;">
<td style="padding:10px 16px;">
<span style="display:inline-block;width:12px;height:12px;border-radius:50%;background:#ff5f56;margin-right:6px;"></span>
<span style="display:inline-block;width:12px;height:12px;border-radius:50%;background:#ffbd2e;margin-right:6px;"></span>
<span style="display:inline-block;width:12px;height:12px;border-radius:50%;background:#27c93f;margin-right:6px;"></span>
<span style="color:#8b949e;font-size:12px;margin-left:8px;">README.md — email-dispatcher</span>
<span style="float:right;font-size:11px;padding:2px 8px;border-radius:20px;background:#00acd722;color:#00acd7;border:1px solid #00acd7;">Go 1.26.3</span>
</td>
</tr>
<tr><td style="padding:28px 32px;">

<!-- Hero -->
<div style="margin-bottom:24px;">
<h1 style="font-size:24px;font-weight:600;color:#e6edf3;display:flex;align-items:center;gap:10px;margin:0 0 8px;">
📨 email-dispatcher
</h1>
<p style="font-size:14px;color:#8b949e;line-height:1.6;margin:0 0 12px;">
A Go backend that reads recipient data from MongoDB and dispatches personalized emails concurrently using a producer-consumer pattern with goroutines.
</p>
<span style="font-size:11px;padding:3px 10px;border-radius:20px;background:#00acd711;color:#00acd7;border:1px solid #00acd7;margin-right:6px;">Go</span>
<span style="font-size:11px;padding:3px 10px;border-radius:20px;background:#4aab4411;color:#4aab44;border:1px solid #4aab44;margin-right:6px;">MongoDB</span>
<span style="font-size:11px;padding:3px 10px;border-radius:20px;background:#f0860011;color:#d97706;border:1px solid #d97706;margin-right:6px;">SMTP</span>
<span style="font-size:11px;padding:3px 10px;border-radius:20px;background:#8b5cf611;color:#8b5cf6;border:1px solid #8b5cf6;">Concurrent</span>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- How it works -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 14px;">⚡ How it works</p>
<table width="100%" style="border-collapse:collapse;text-align:center;">
<tr>
<td style="padding:0 4px;">
<div style="width:36px;height:36px;border-radius:10px;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;margin:0 auto 6px;font-size:16px;">🗄️</div>
<div style="font-size:12px;font-weight:600;color:#e6edf3;">MongoDB</div>
<div style="font-size:11px;color:#8b949e;">Load contacts</div>
</td>
<td style="color:#30363d;font-size:18px;padding-bottom:14px;">──</td>
<td style="padding:0 4px;">
<div style="width:36px;height:36px;border-radius:10px;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;margin:0 auto 6px;font-size:16px;">⇄</div>
<div style="font-size:12px;font-weight:600;color:#e6edf3;">Channel</div>
<div style="font-size:11px;color:#8b949e;">Push to channel</div>
</td>
<td style="color:#30363d;font-size:18px;padding-bottom:14px;">──</td>
<td style="padding:0 4px;">
<div style="width:36px;height:36px;border-radius:10px;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;margin:0 auto 6px;font-size:16px;">⚙️</div>
<div style="font-size:12px;font-weight:600;color:#e6edf3;">Workers</div>
<div style="font-size:11px;color:#8b949e;">N goroutines</div>
</td>
<td style="color:#30363d;font-size:18px;padding-bottom:14px;">──</td>
<td style="padding:0 4px;">
<div style="width:36px;height:36px;border-radius:10px;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;margin:0 auto 6px;font-size:16px;">📤</div>
<div style="font-size:12px;font-weight:600;color:#e6edf3;">SMTP</div>
<div style="font-size:11px;color:#8b949e;">Send email</div>
</td>
<td style="color:#30363d;font-size:18px;padding-bottom:14px;">──</td>
<td style="padding:0 4px;">
<div style="width:36px;height:36px;border-radius:10px;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;margin:0 auto 6px;font-size:16px;">✅</div>
<div style="font-size:12px;font-weight:600;color:#e6edf3;">WaitGroup</div>
<div style="font-size:11px;color:#8b949e;">All done</div>
</td>
</tr>
</table>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Project Structure -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">📁 Project structure</p>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:14px 16px;font-size:12.5px;line-height:1.9;color:#8b949e;">
<span style="color:#00c896;font-weight:600;">email-dispatcher/</span><br>
├── <span style="color:#00acd7;">main.go</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span style="opacity:.5;"># entry point, flag parsing, wires producer &amp; consumers</span><br>
├── <span style="color:#00acd7;">producer.go</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span style="opacity:.5;"># reads recipients from MongoDB → channel</span><br>
├── <span style="color:#00acd7;">consumer.go</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span style="opacity:.5;"># worker goroutines that send emails via SMTP</span><br>
├── <span style="color:#00acd7;">db.go</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span style="opacity:.5;"># MongoDB connection, index creation, recipient loader</span><br>
├── <span style="color:#e6edf3;">email.tmpl</span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<span style="opacity:.5;"># Go HTML template for email body</span><br>
├── <span style="color:#e6edf3;">name_emails_200_records.csv</span>&nbsp;&nbsp;<span style="opacity:.5;"># sample CSV for seeding</span><br>
├── <span style="color:#e6edf3;">go.mod</span><br>
└── <span style="color:#e6edf3;">go.sum</span>
</div>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Tech Stack -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">🛠️ Tech stack</p>
<table width="100%" style="border-collapse:collapse;">
<tr>
<td width="20%" style="padding:4px 8px 4px 0;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:10px 12px;">
<div style="font-size:10px;color:#8b949e;margin-bottom:3px;">Language</div>
<div style="font-size:13px;font-weight:600;color:#00acd7;">Go 1.26.3</div>
</div>
</td>
<td width="20%" style="padding:4px 8px;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:10px 12px;">
<div style="font-size:10px;color:#8b949e;margin-bottom:3px;">Database</div>
<div style="font-size:13px;font-weight:600;color:#4aab44;">MongoDB</div>
</div>
</td>
<td width="20%" style="padding:4px 8px;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:10px 12px;">
<div style="font-size:10px;color:#8b949e;margin-bottom:3px;">Email</div>
<div style="font-size:13px;font-weight:600;color:#e6edf3;">net/smtp</div>
</div>
</td>
<td width="20%" style="padding:4px 8px;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:10px 12px;">
<div style="font-size:10px;color:#8b949e;margin-bottom:3px;">Templating</div>
<div style="font-size:13px;font-weight:600;color:#e6edf3;">html/template</div>
</div>
</td>
<td width="20%" style="padding:4px 8px 4px 0;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:10px 12px;">
<div style="font-size:10px;color:#8b949e;margin-bottom:3px;">Concurrency</div>
<div style="font-size:13px;font-weight:600;color:#8b5cf6;">goroutines</div>
</div>
</td>
</tr>
</table>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Prerequisites -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">✅ Prerequisites</p>
<div style="font-size:13px;color:#e6edf3;line-height:2;font-family:'Courier New',monospace;">
&nbsp;&nbsp;🔹 Go 1.26.3+<br>
&nbsp;&nbsp;🔹 MongoDB running at <code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:1px 6px;color:#00acd7;">localhost:27017</code><br>
&nbsp;&nbsp;🔹 SMTP server at <code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:1px 6px;color:#00acd7;">localhost:1025</code> — e.g. MailHog for local dev
</div>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Setup -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 14px;">⚙️ Setup</p>

<div style="display:flex;gap:12px;margin-bottom:12px;">
<div style="width:24px;height:24px;border-radius:50%;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;font-size:11px;font-weight:600;color:#00c896;flex-shrink:0;">1</div>
<div style="flex:1;">
<div style="font-size:13px;font-weight:600;color:#e6edf3;margin-bottom:6px;">Clone the repository</div>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:10px 14px;font-size:12px;line-height:1.7;"><span style="color:#00c896;">git clone</span> <span style="color:#e0957a;">https://github.com/saurabhisane/email-dispatcher.git</span><br><span style="color:#00c896;">cd</span> email-dispatcher</div>
</div>
</div>

<div style="display:flex;gap:12px;margin-bottom:12px;">
<div style="width:24px;height:24px;border-radius:50%;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;font-size:11px;font-weight:600;color:#00c896;flex-shrink:0;">2</div>
<div style="flex:1;">
<div style="font-size:13px;font-weight:600;color:#e6edf3;margin-bottom:6px;">Install dependencies</div>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:10px 14px;font-size:12px;"><span style="color:#00c896;">go mod tidy</span></div>
</div>
</div>

<div style="display:flex;gap:12px;margin-bottom:12px;">
<div style="width:24px;height:24px;border-radius:50%;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;font-size:11px;font-weight:600;color:#00c896;flex-shrink:0;">3</div>
<div style="flex:1;">
<div style="font-size:13px;font-weight:600;color:#e6edf3;margin-bottom:6px;">Seed MongoDB</div>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:10px 14px;font-size:12px;line-height:1.7;color:#8b949e;"><span style="opacity:.6;"># Import name_emails_200_records.csv into the contacts collection</span><br><span style="opacity:.6;"># Each document requires: name, email fields</span><br><span style="opacity:.6;"># Database: email_dispatcher  |  Collection: contacts</span></div>
</div>
</div>

<div style="display:flex;gap:12px;">
<div style="width:24px;height:24px;border-radius:50%;background:#00c89618;border:1px solid #00c89655;display:flex;align-items:center;justify-content:center;font-size:11px;font-weight:600;color:#00c896;flex-shrink:0;">4</div>
<div style="flex:1;">
<div style="font-size:13px;font-weight:600;color:#e6edf3;margin-bottom:6px;">Run the application</div>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:10px 14px;font-size:12px;line-height:1.7;"><span style="color:#00c896;">go run</span> .<br><br><span style="color:#8b949e;opacity:.6;"># with custom flags:</span><br><span style="color:#00c896;">go run</span> . <span style="color:#f0a500;">-mongoURI</span>=<span style="color:#e0957a;">"mongodb://localhost:27017"</span> <span style="color:#f0a500;">-dbName</span>=<span style="color:#e0957a;">"email_dispatcher"</span> <span style="color:#f0a500;">-workerCount</span>=10</div>
</div>
</div>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- CLI Flags -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">🚩 CLI flags</p>
<table width="100%" style="border-collapse:collapse;font-size:12.5px;">
<tr style="border-bottom:1px solid #2d333b;">
<th style="text-align:left;font-size:11px;text-transform:uppercase;letter-spacing:.07em;color:#8b949e;padding:0 12px 8px 0;font-weight:500;">Flag</th>
<th style="text-align:left;font-size:11px;text-transform:uppercase;letter-spacing:.07em;color:#8b949e;padding:0 12px 8px;font-weight:500;">Default</th>
<th style="text-align:left;font-size:11px;text-transform:uppercase;letter-spacing:.07em;color:#8b949e;padding:0 0 8px;font-weight:500;">Description</th>
</tr>
<tr style="border-bottom:1px solid #2d333b;">
<td style="padding:9px 12px 9px 0;"><code style="color:#00acd7;background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:12px;">-mongoURI</code></td>
<td style="padding:9px 12px;"><code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:11px;color:#8b949e;">mongodb://localhost:27017</code></td>
<td style="padding:9px 0;color:#e6edf3;">MongoDB connection URI</td>
</tr>
<tr style="border-bottom:1px solid #2d333b;">
<td style="padding:9px 12px 9px 0;"><code style="color:#00acd7;background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:12px;">-dbName</code></td>
<td style="padding:9px 12px;"><code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:11px;color:#8b949e;">email_dispatcher</code></td>
<td style="padding:9px 0;color:#e6edf3;">MongoDB database name</td>
</tr>
<tr>
<td style="padding:9px 12px 9px 0;"><code style="color:#00acd7;background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:12px;">-workerCount</code></td>
<td style="padding:9px 12px;"><code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:2px 6px;font-size:11px;color:#8b949e;">5</code></td>
<td style="padding:9px 0;color:#e6edf3;">Number of concurrent email worker goroutines</td>
</tr>
</table>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- MongoDB Schema -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">🗄️ MongoDB schema</p>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:12px 16px;font-size:12px;line-height:1.8;margin-bottom:10px;">
{<br>
&nbsp;&nbsp;<span style="color:#f0a500;">"name"</span>:&nbsp;&nbsp;<span style="color:#e0957a;">"John Doe"</span>,<br>
&nbsp;&nbsp;<span style="color:#f0a500;">"email"</span>: <span style="color:#e0957a;">"johndoe@example.com"</span><br>
}
</div>
<div style="font-size:12px;color:#8b949e;display:flex;align-items:center;gap:6px;">
ℹ️ Collection: <code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:1px 5px;color:#e6edf3;">contacts</code> — a unique index on <code style="background:#161b22;border:1px solid #2d333b;border-radius:4px;padding:1px 5px;color:#e6edf3;">email</code> is created automatically at startup.
</div>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Email Template -->
<div style="margin-bottom:20px;">
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">📧 Email template</p>
<table width="100%" style="border-collapse:collapse;">
<tr>
<td width="50%" style="padding-right:8px;vertical-align:top;">
<div style="font-size:10px;color:#8b949e;text-transform:uppercase;letter-spacing:.08em;margin-bottom:6px;">email.tmpl</div>
<div style="background:#161b22;border:1px solid #2d333b;border-left:3px solid #00c896;border-radius:6px;padding:12px 14px;font-size:12px;line-height:1.8;color:#e6edf3;">
Subject: Hello, <span style="color:#f0a500;">{{.Name}}</span><br><br>
hi <span style="color:#f0a500;">{{.Name}}</span><br><br>
This is a test email to check<br>
the functionality of our email<br>
dispatcher.<br><br>
Thank you,<br>
Saurabh Isane
</div>
</td>
<td width="50%" style="padding-left:8px;vertical-align:top;">
<div style="font-size:10px;color:#8b949e;text-transform:uppercase;letter-spacing:.08em;margin-bottom:6px;">rendered output</div>
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:14px 16px;font-size:13px;line-height:1.8;color:#e6edf3;font-family:sans-serif;">
<div style="font-size:10px;color:#8b949e;text-transform:uppercase;letter-spacing:.06em;margin-bottom:8px;">Subject: Hello, <span style="color:#00c896;font-weight:600;">Alex</span></div>
hi <span style="color:#00c896;font-weight:600;">Alex</span><br><br>
This is a test email to check the functionality of our email dispatcher.<br><br>
Thank you,<br>
Saurabh Isane
</div>
</td>
</tr>
</table>
</div>

<hr style="border:none;border-top:1px solid #2d333b;margin:20px 0;">

<!-- Dependencies -->
<div>
<p style="font-size:11px;font-weight:600;text-transform:uppercase;letter-spacing:.08em;color:#8b949e;margin:0 0 12px;">📦 Dependencies</p>
<table width="100%" style="border-collapse:collapse;">
<tr>
<td width="50%" style="padding:4px 6px 4px 0;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:8px 12px;font-size:12px;display:flex;align-items:center;gap:8px;">
<span style="width:6px;height:6px;border-radius:50%;background:#4aab44;display:inline-block;"></span>
<span style="color:#e6edf3;">go.mongodb.org/mongo-driver</span>
<span style="margin-left:auto;font-size:11px;color:#8b949e;">v1.17.9</span>
</div>
</td>
<td width="50%" style="padding:4px 0 4px 6px;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:8px 12px;font-size:12px;display:flex;align-items:center;gap:8px;">
<span style="width:6px;height:6px;border-radius:50%;background:#00acd7;display:inline-block;"></span>
<span style="color:#e6edf3;">golang.org/x/crypto</span>
<span style="margin-left:auto;font-size:11px;color:#8b949e;">indirect</span>
</div>
</td>
</tr>
<tr>
<td style="padding:4px 6px 0 0;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:8px 12px;font-size:12px;display:flex;align-items:center;gap:8px;">
<span style="width:6px;height:6px;border-radius:50%;background:#00acd7;display:inline-block;"></span>
<span style="color:#e6edf3;">golang.org/x/sync</span>
<span style="margin-left:auto;font-size:11px;color:#8b949e;">indirect</span>
</div>
</td>
<td style="padding:4px 0 0 6px;">
<div style="background:#161b22;border:1px solid #2d333b;border-radius:6px;padding:8px 12px;font-size:12px;display:flex;align-items:center;gap:8px;">
<span style="width:6px;height:6px;border-radius:50%;background:#00acd7;display:inline-block;"></span>
<span style="color:#e6edf3;">golang.org/x/text</span>
<span style="margin-left:auto;font-size:11px;color:#8b949e;">indirect</span>
</div>
</td>
</tr>
</table>
</div>

</td></tr>
</table>

</div>
