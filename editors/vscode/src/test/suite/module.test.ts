import * as assert from 'assert';
import * as vscode from 'vscode';
import * as path from 'path';

suite('GoML Module Tests', () => {
    const mainGoml = path.resolve(__dirname, '../../../../../gomlc/testdata/module/project001/main.goml');

    async function openModule(): Promise<vscode.TextDocument> {
        const document = await vscode.workspace.openTextDocument(mainGoml);
        await vscode.window.showTextDocument(document);
        return document;
    }

    test('Extension should be present', () => {
        assert.ok(vscode.extensions.getExtension('goml.goml'));
    });

    test('Should activate on .goml files', async () => {
        const ext = vscode.extensions.getExtension('goml.goml');
        assert.ok(ext, 'Extension should be present');
        await openModule();
        await new Promise(resolve => setTimeout(resolve, 2000));
        assert.ok(ext.isActive, 'Extension should be active after opening .goml file');
    });

    test('Hover should work on module files', async function () {
        this.timeout(30000);
        const doc = await openModule();
        await new Promise(resolve => setTimeout(resolve, 3000));

        const match = /fn\s+(main)/.exec(doc.getText());
        assert.ok(match, 'Module fixture should declare main');
        const pos = doc.positionAt(match.index + match[0].lastIndexOf('main') + 2);
        const hovers = await vscode.commands.executeCommand<vscode.Hover[]>(
            'vscode.executeHoverProvider',
            doc.uri,
            pos
        );

        assert.ok(hovers && hovers.length > 0, 'Hover should describe the main function');
    });

    test('Diagnostics should work for module files', async function () {
        this.timeout(30000);
        const doc = await openModule();
        await new Promise(resolve => setTimeout(resolve, 5000));

        const diagnostics = vscode.languages.getDiagnostics(doc.uri);
        assert.strictEqual(diagnostics.length, 0, 'Valid module file should have no diagnostics');
    });
});
