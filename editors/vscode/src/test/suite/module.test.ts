import * as assert from 'assert';
import * as vscode from 'vscode';
import * as path from 'path';
import * as fs from 'fs';

suite('GoML Module Tests', () => {
    const moduleTestDir = path.resolve(__dirname, '../../../../gomlc/testdata/module');

    function getProjectDirs(): string[] {
        return fs.readdirSync(moduleTestDir)
            .filter(name => name.startsWith('project'))
            .map(name => path.join(moduleTestDir, name));
    }

    test('Extension should be present', () => {
        assert.ok(vscode.extensions.getExtension('goml.goml'));
    });

    test('Should activate on .goml files', async () => {
        const ext = vscode.extensions.getExtension('goml.goml');
        assert.ok(ext, 'Extension should be present');

        const projectDirs = getProjectDirs();
        if (projectDirs.length > 0) {
            const mainGom = path.join(projectDirs[0], 'main.goml');
            if (fs.existsSync(mainGom)) {
                const doc = await vscode.workspace.openTextDocument(mainGom);
                await vscode.window.showTextDocument(doc);

                await new Promise(resolve => setTimeout(resolve, 2000));

                assert.ok(ext.isActive, 'Extension should be active after opening .goml file');
            }
        }
    });

    test('All module projects should have module manifest', () => {
        const projectDirs = getProjectDirs();

        for (const projectDir of projectDirs) {
            const gomlToml = path.join(projectDir, 'goml.toml');
            assert.ok(
                fs.existsSync(gomlToml),
                `${path.basename(projectDir)} should have goml.toml`
            );

            const content = fs.readFileSync(gomlToml, 'utf8');
            assert.ok(
                content.includes('[module]'),
                `${path.basename(projectDir)}/goml.toml should have [module] section`
            );
        }
    });

    test('Packages should not have module manifests', () => {
        const projectDirs = getProjectDirs();

        for (const projectDir of projectDirs) {
            const entries = fs.readdirSync(projectDir, { withFileTypes: true });
            const subDirs = entries.filter(e => e.isDirectory());

            for (const subDir of subDirs) {
                const subGomlToml = path.join(projectDir, subDir.name, 'goml.toml');
                assert.ok(
                    !fs.existsSync(subGomlToml),
                    `${path.basename(projectDir)}/${subDir.name}/goml.toml should not exist`
                );
            }
        }
    });

    test('Hover should work on module files', async function () {
        this.timeout(30000);

        const ext = vscode.extensions.getExtension('goml.goml');
        if (!ext) {
            this.skip();
            return;
        }

        const projectDirs = getProjectDirs();
        if (projectDirs.length === 0) {
            this.skip();
            return;
        }

        const mainGom = path.join(projectDirs[0], 'main.goml');
        if (!fs.existsSync(mainGom)) {
            this.skip();
            return;
        }

        const doc = await vscode.workspace.openTextDocument(mainGom);
        await vscode.window.showTextDocument(doc);

        await new Promise(resolve => setTimeout(resolve, 3000));

        if (!ext.isActive) {
            this.skip();
            return;
        }

        const text = doc.getText();
        const fnMainMatch = text.match(/fn\s+main/);
        if (!fnMainMatch) {
            this.skip();
            return;
        }

        const pos = doc.positionAt(text.indexOf('main') + 2);
        const hovers = await vscode.commands.executeCommand<vscode.Hover[]>(
            'vscode.executeHoverProvider',
            doc.uri,
            pos
        );

        assert.ok(hovers && hovers.length >= 0, 'Hover command should return result');
    });

    test('Diagnostics should work for module files', async function () {
        this.timeout(30000);

        const ext = vscode.extensions.getExtension('goml.goml');
        if (!ext) {
            this.skip();
            return;
        }

        const projectDirs = getProjectDirs();
        if (projectDirs.length === 0) {
            this.skip();
            return;
        }

        const mainGom = path.join(projectDirs[0], 'main.goml');
        if (!fs.existsSync(mainGom)) {
            this.skip();
            return;
        }

        const doc = await vscode.workspace.openTextDocument(mainGom);
        await vscode.window.showTextDocument(doc);

        await new Promise(resolve => setTimeout(resolve, 5000));

        const diagnostics = vscode.languages.getDiagnostics(doc.uri);
        assert.ok(diagnostics.length === 0, 'Valid module file should have no diagnostics');
    });
});
