import * as assert from 'assert';
import * as fs from 'fs';
import * as path from 'path';
import { Registry, INITIAL, IGrammar, parseRawGrammar } from 'vscode-textmate';
import { loadWASM, OnigScanner, OnigString } from 'vscode-oniguruma';

interface GrammarPattern {
    include?: string;
    match?: string;
}

interface GrammarRule {
    patterns: GrammarPattern[];
}

interface Grammar {
    patterns: GrammarPattern[];
    repository: Record<string, GrammarRule>;
}

suite('GoML Syntax Tests', () => {
    const grammarPath = path.resolve(__dirname, '../../../syntaxes/goml.tmLanguage.json');
    const grammar = JSON.parse(fs.readFileSync(grammarPath, 'utf8')) as Grammar;

    function regex(rule: string, pattern: number): RegExp {
        const value = grammar.repository[rule].patterns[pattern].match;
        assert.ok(value, `Expected ${rule} pattern ${pattern} to define a match`);
        return new RegExp(value);
    }

    test('Loop labels take precedence over character literals', () => {
        const labels = grammar.patterns.findIndex(pattern => pattern.include === '#labels');
        const chars = grammar.patterns.findIndex(pattern => pattern.include === '#chars');
        assert.ok(labels >= 0);
        assert.ok(chars >= 0);
        assert.ok(labels < chars);

        const definition = regex('labels', 0);
        const reference = regex('labels', 1);
        assert.ok(definition.test("'outer: loop"));
        assert.ok(reference.test("break 'outer;"));
        assert.ok(reference.test("continue 'outer;"));
        assert.ok(!definition.test("'a'"));
        assert.ok(!definition.test("'ab'"));
    });

    test('Public imports and dyn associated bindings are recognized', () => {
        assert.ok(regex('declarations', 0).test('pub use client::Request as HttpRequest;'));
        assert.ok(regex('declarations', 0).test('use std::task;'));
        assert.ok(regex('associated-type-bindings', 0).test('dyn Iterator[Item = int]'));
    });

    test('Task keywords remain contextual', () => {
        const contextual = regex('keywords', 0);
        assert.ok(contextual.test('scope {'));
        assert.ok(contextual.test('spawn |cancel| work(cancel)'));
        assert.ok(!contextual.test('let scope = 3;'));
        assert.ok(!contextual.test('spawn(scope)'));
    });

    test('Select keywords remain contextual', () => {
        const expressions = regex('keywords', 0);
        const arms = regex('keywords', 1);
        assert.ok(expressions.test('select {'));
        assert.ok(expressions.test('select priority {'));
        assert.ok(expressions.test('priority {'));
        assert.ok(arms.test('recv(input) as value => value'));
        assert.ok(arms.test('recv(input) when enabled as value => value'));
        assert.ok(arms.test('recv(input) when enabled match { Some(value) => value }'));
        assert.ok(arms.test('send(output, select(1)) => 2'));
        assert.ok(arms.test('send(output, value) when enabled => 2'));
        assert.ok(arms.test('default => 3'));
        assert.ok(!expressions.test('select(value)'));
        assert.ok(!arms.test('channel.recv()'));
        assert.ok(!arms.test('channel.send(value)'));
        assert.ok(!arms.test('let default = 3;'));
        assert.ok(!expressions.test('priority(value)'));
        assert.ok(!arms.test('let when = true;'));
    });
});

suite('GoML TextMate Tokenization', () => {
  let registry: Registry;
  let grammar: IGrammar;

  suiteSetup(async () => {
    const wasm = fs.readFileSync(require.resolve('vscode-oniguruma/release/onig.wasm'));
    await loadWASM(wasm.buffer.slice(wasm.byteOffset, wasm.byteOffset + wasm.byteLength));
    registry = new Registry({
      onigLib: Promise.resolve({
        createOnigScanner: patterns => new OnigScanner(patterns),
        createOnigString: value => new OnigString(value),
      }),
      loadGrammar: async () => {
        const filename = path.resolve(__dirname, '../../../syntaxes/goml.tmLanguage.json');
        return parseRawGrammar(fs.readFileSync(filename, 'utf8'), filename);
      },
    });
    const loaded = await registry.loadGrammar('source.goml');
    assert.ok(loaded);
    grammar = loaded;
  });

  suiteTeardown(() => registry.dispose());

  function scopesAt(source: string, text: string): string[] {
    const offset = source.indexOf(text);
    assert.ok(offset >= 0, text);
    const token = grammar.tokenizeLine(source, INITIAL).tokens.find(
      token => token.startIndex <= offset && token.endIndex > offset,
    );
    assert.ok(token, text);
    return token.scopes;
  }

  test('Literal braces and adjacent expression braces close correctly', () => {
    for (const literal of [
      'f"{{{value}}}"',
      'f"{value}}}"',
      'f"{ {value}}"',
      'f"{ { {value}}}"',
      'f"{if true {1} else {2}}"',
      'f"{r"}"}"',
      'f"{f"{value}"}"',
    ]) {
      const source = `let text = ${literal}; let after = 1;`;
      const result = grammar.tokenizeLine(source, INITIAL);
      assert.ok(result.ruleStack.equals(grammar.tokenizeLine('', INITIAL).ruleStack), literal);
      assert.deepStrictEqual(scopesAt(source, 'after'), ['source.goml'], literal);
      const next = grammar.tokenizeLine('let next = 2;', result.ruleStack);
      assert.deepStrictEqual(next.tokens[0].scopes, ['source.goml', 'keyword.declaration.goml']);
    }
    const source = 'f"{{{value}}}"';
    assert.ok(scopesAt(source, '{{').includes('constant.character.escape.goml'));
    assert.ok(scopesAt(source, 'value').includes('meta.interpolation.goml'));
    assert.ok(scopesAt('f"{{text}}"', '}}').includes('constant.character.escape.goml'));
    assert.ok(scopesAt('f"\\q"', '\\q').includes('invalid.illegal.escape.goml'));
  });

  test('Generic calls retain function scopes with nested type arguments', () => {
    for (const source of [
      'identity::[i32](1)',
      'pkg::identity::[Vec[Option[i32]]](value)',
      'value.convert::[string](fallback)',
      'Box::[i32]::convert::[string](value)',
      'identity :: [i32](1)',
    ]) {
      const name = source.includes('identity') ? 'identity' : 'convert';
      assert.ok(scopesAt(source, name).includes('entity.name.function.goml'), source);
    }
    assert.ok(scopesAt('pkg::identity(1)', 'pkg').includes('entity.name.namespace.goml'));
    assert.ok(scopesAt('Box::new()', 'Box').includes('entity.name.type.goml'));
    assert.ok(!scopesAt('values[index](1)', 'values').includes('entity.name.function.goml'));
  });

  test('Constants and statics are values while declarations remain types', () => {
    assert.ok(scopesAt('pub const MAX: isize = 10;', 'MAX').includes('constant.other.goml'));
    assert.ok(scopesAt('static cache: OnceCell[string];', 'cache').includes('variable.other.goml'));
    for (const declaration of ['struct', 'enum', 'trait', 'type']) {
      assert.ok(scopesAt(`${declaration} Example`, 'Example').includes('entity.name.type.goml'));
    }
  });

  test('Primitive types are distinct and library types use ordinary type scopes', () => {
    for (const name of ['bool', 'isize', 'string', 'byte', 'never']) {
      assert.ok(scopesAt(`fn value() -> ${name}`, name).includes('storage.type.primitive.goml'));
    }
    for (const name of ['Vec', 'Option', 'HashMap', 'Bytes', 'OnceCell', 'FrozenVec', 'Custom']) {
      assert.deepStrictEqual(scopesAt(`let value: ${name};`, name), ['source.goml', 'entity.name.type.goml']);
    }
  });
});
