<script lang="ts">
	let formData = {
		nombre_completo: '',
		email: '',
		telefono: '',
		tipo_producto: '',
		modelo_producto: '',
		numero_serie: '',
		categoria_problema: '',
		descripcion_problema: '',
		sistema_operativo: '',
		fecha_compra: '',
		lugar_compra: '',
		metodo_contacto_preferido: 'email',
		horario_contacto: '',
		acepta_terminos: false,
		acepta_marketing: false
	};

	let errors: Record<string, string> = {};
	let isSubmitting = false;
	let submitSuccess = false;
	let submitError = '';

	const tiposProducto = [
		'Ratón',
		'Teclado',
		'Auriculares',
		'Webcam',
		'Altavoces',
		'Controlador de juegos',
		'Otro'
	];

	const categoriasProblema = [
		'Problema de conexión',
		'No enciende',
		'Botones no funcionan',
		'Problema de software/drivers',
		'Problema de batería',
		'Problema de audio',
		'Problema de video',
		'Garantía',
		'Otro'
	];

	const sistemasOperativos = [
		'Windows 11',
		'Windows 10',
		'macOS Sonoma',
		'macOS Ventura',
		'Linux',
		'ChromeOS',
		'Otro'
	];

	const horariosContacto = [
		'Mañana (9:00 - 12:00)',
		'Tarde (12:00 - 18:00)',
		'Noche (18:00 - 21:00)',
		'Cualquier horario'
	];

	function validateForm(): boolean {
		errors = {};

		if (!formData.nombre_completo.trim()) {
			errors.nombre_completo = 'El nombre completo es requerido';
		}

		if (!formData.email.trim()) {
			errors.email = 'El email es requerido';
		} else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
			errors.email = 'El email no es válido';
		}

		if (!formData.tipo_producto) {
			errors.tipo_producto = 'Selecciona el tipo de producto';
		}

		if (!formData.categoria_problema) {
			errors.categoria_problema = 'Selecciona la categoría del problema';
		}

		if (!formData.descripcion_problema.trim()) {
			errors.descripcion_problema = 'La descripción del problema es requerida';
		} else if (formData.descripcion_problema.trim().length < 20) {
			errors.descripcion_problema = 'La descripción debe tener al menos 20 caracteres';
		}

		if (!formData.acepta_terminos) {
			errors.acepta_terminos = 'Debes aceptar los términos y condiciones';
		}

		return Object.keys(errors).length === 0;
	}

	async function handleSubmit(e: Event) {
		e.preventDefault();
		submitError = '';
		submitSuccess = false;

		if (!validateForm()) {
			return;
		}

		isSubmitting = true;

		try {
			const response = await fetch('/api/v1/formularios', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify(formData)
			});

			const data = await response.json();

			if (response.ok && !data.error) {
				submitSuccess = true;
				// Resetear formulario
				formData = {
					nombre_completo: '',
					email: '',
					telefono: '',
					tipo_producto: '',
					modelo_producto: '',
					numero_serie: '',
					categoria_problema: '',
					descripcion_problema: '',
					sistema_operativo: '',
					fecha_compra: '',
					lugar_compra: '',
					metodo_contacto_preferido: 'email',
					horario_contacto: '',
					acepta_terminos: false,
					acepta_marketing: false
				};
				// Scroll al top
				window.scrollTo({ top: 0, behavior: 'smooth' });
			} else {
				submitError = data.mensaje || 'Error al enviar el formulario';
			}
		} catch (error) {
			submitError = 'Error de conexión. Por favor, intenta nuevamente.';
			console.error('Error:', error);
		} finally {
			isSubmitting = false;
		}
	}

	function resetForm() {
		submitSuccess = false;
		submitError = '';
	}
</script>

<svelte:head>
	<title>Soporte Técnico Logitech - Formulario de Contacto</title>
	<meta name="description" content="Formulario de soporte técnico para productos Logitech" />
</svelte:head>

<div class="min-h-screen bg-gradient-to-br from-logitech-darkblue via-gray-900 to-logitech-darkblue py-12 px-4 sm:px-6 lg:px-8">
	<div class="max-w-4xl mx-auto">
		<!-- Header -->
		<div class="text-center mb-12">
			<h1 class="text-4xl md:text-5xl font-bold text-white mb-4">
				Soporte Técnico Logitech
			</h1>
			<p class="text-xl text-gray-300">
				Estamos aquí para ayudarte con tu producto Logitech
			</p>
		</div>

		<!-- Success Message -->
		{#if submitSuccess}
			<div class="bg-green-500 text-white p-6 rounded-lg shadow-xl mb-8 animate-pulse">
				<div class="flex items-center">
					<svg class="w-8 h-8 mr-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path>
					</svg>
					<div>
						<h3 class="text-xl font-bold mb-1">¡Formulario enviado exitosamente!</h3>
						<p>Nuestro equipo de soporte se pondrá en contacto contigo pronto.</p>
					</div>
				</div>
				<button on:click={resetForm} class="mt-4 underline hover:no-underline">
					Enviar otro formulario
				</button>
			</div>
		{/if}

		<!-- Error Message -->
		{#if submitError}
			<div class="bg-red-500 text-white p-6 rounded-lg shadow-xl mb-8">
				<div class="flex items-center">
					<svg class="w-8 h-8 mr-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
					</svg>
					<div>
						<h3 class="text-xl font-bold mb-1">Error al enviar el formulario</h3>
						<p>{submitError}</p>
					</div>
				</div>
			</div>
		{/if}

		<!-- Form -->
		<div class="bg-white rounded-2xl shadow-2xl p-8 md:p-12">
			<form on:submit={handleSubmit} class="space-y-8">
				<!-- Información Personal -->
				<div class="border-b border-gray-200 pb-8">
					<h2 class="text-2xl font-bold text-gray-900 mb-6 flex items-center">
						<span class="bg-logitech-blue text-white w-8 h-8 rounded-full flex items-center justify-center mr-3">1</span>
						Información Personal
					</h2>
					
					<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
						<div class="md:col-span-2">
							<label for="nombre_completo" class="form-label">
								Nombre Completo <span class="text-red-500">*</span>
							</label>
							<input
								type="text"
								id="nombre_completo"
								bind:value={formData.nombre_completo}
								class="form-input"
								class:border-red-500={errors.nombre_completo}
								placeholder="Juan Pérez García"
							/>
							{#if errors.nombre_completo}
								<p class="form-error">{errors.nombre_completo}</p>
							{/if}
						</div>

						<div>
							<label for="email" class="form-label">
								Email <span class="text-red-500">*</span>
							</label>
							<input
								type="email"
								id="email"
								bind:value={formData.email}
								class="form-input"
								class:border-red-500={errors.email}
								placeholder="juan.perez@ejemplo.com"
							/>
							{#if errors.email}
								<p class="form-error">{errors.email}</p>
							{/if}
						</div>

						<div>
							<label for="telefono" class="form-label">
								Teléfono
							</label>
							<input
								type="tel"
								id="telefono"
								bind:value={formData.telefono}
								class="form-input"
								placeholder="+56 9 1234 5678"
							/>
						</div>
					</div>
				</div>

				<!-- Información del Producto -->
				<div class="border-b border-gray-200 pb-8">
					<h2 class="text-2xl font-bold text-gray-900 mb-6 flex items-center">
						<span class="bg-logitech-blue text-white w-8 h-8 rounded-full flex items-center justify-center mr-3">2</span>
						Información del Producto
					</h2>
					
					<div class="grid grid-cols-1 md:grid-cols-2 gap-6">
						<div>
							<label for="tipo_producto" class="form-label">
								Tipo de Producto <span class="text-red-500">*</span>
							</label>
							<select
								id="tipo_producto"
								bind:value={formData.tipo_producto}
								class="form-input"
								class:border-red-500={errors.tipo_producto}
							>
								<option value="">Selecciona un tipo</option>
								{#each tiposProducto as tipo}
									<option value={tipo}>{tipo}</option>
								{/each}
							</select>
							{#if errors.tipo_producto}
								<p class="form-error">{errors.tipo_producto}</p>
							{/if}
						</div>

						<div>
							<label for="modelo_producto" class="form-label">
								Modelo del Producto
							</label>
							<input
								type="text"
								id="modelo_producto"
								bind:value={formData.modelo_producto}
								class="form-input"
								placeholder="Ej: MX Master 3S"
							/>
						</div>

						<div>
							<label for="numero_serie" class="form-label">
								Número de Serie
							</label>
							<input
								type="text"
								id="numero_serie"
								bind:value={formData.numero_serie}
								class="form-input"
								placeholder="Ej: SN123456789"
							/>
						</div>

						<div>
							<label for="fecha_compra" class="form-label">
								Fecha de Compra
							</label>
							<input
								type="date"
								id="fecha_compra"
								bind:value={formData.fecha_compra}
								class="form-input"
							/>
						</div>

						<div class="md:col-span-2">
							<label for="lugar_compra" class="form-label">
								Lugar de Compra
							</label>
							<input
								type="text"
								id="lugar_compra"
								bind:value={formData.lugar_compra}
								class="form-input"
								placeholder="Ej: Amazon, Falabella, Tienda oficial"
							/>
						</div>
					</div>
				</div>

				<!-- Detalles del Problema -->
				<div class="border-b border-gray-200 pb-8">
					<h2 class="text-2xl font-bold text-gray-900 mb-6 flex items-center">
						<span class="bg-logitech-blue text-white w-8 h-8 rounded-full flex items-center justify-center mr-3">3</span>
						Detalles del Problema
					</h2>
					
					<div class="space-y-6">
						<div>
							<label for="categoria_problema" class="form-label">
								Categoría del Problema <span class="text-red-500">*</span>
							</label>
							<select
								id="categoria_problema"
								bind:value={formData.categoria_problema}
								class="form-input"
								class:border-red-500={errors.categoria_problema}
							>
								<option value="">Selecciona una categoría</option>
								{#each categoriasProblema as categoria}
									<option value={categoria}>{categoria}</option>
								{/each}
							</select>
							{#if errors.categoria_problema}
								<p class="form-error">{errors.categoria_problema}</p>
							{/if}
						</div>

						<div>
							<label for="descripcion_problema" class="form-label">
								Descripción del Problema <span class="text-red-500">*</span>
							</label>
							<textarea
								id="descripcion_problema"
								bind:value={formData.descripcion_problema}
								rows="6"
								class="form-input resize-none"
								class:border-red-500={errors.descripcion_problema}
								placeholder="Por favor, describe tu problema con el mayor detalle posible. Incluye cuándo comenzó, qué has intentado hacer para solucionarlo, etc."
							></textarea>
							<p class="text-sm text-gray-500 mt-1">
								Mínimo 20 caracteres ({formData.descripcion_problema.length}/20)
							</p>
							{#if errors.descripcion_problema}
								<p class="form-error">{errors.descripcion_problema}</p>
							{/if}
						</div>

						<div>
							<label for="sistema_operativo" class="form-label">
								Sistema Operativo
							</label>
							<select
								id="sistema_operativo"
								bind:value={formData.sistema_operativo}
								class="form-input"
							>
								<option value="">Selecciona tu sistema operativo</option>
								{#each sistemasOperativos as so}
									<option value={so}>{so}</option>
								{/each}
							</select>
						</div>
					</div>
				</div>

				<!-- Preferencias de Contacto -->
				<div class="border-b border-gray-200 pb-8">
					<h2 class="text-2xl font-bold text-gray-900 mb-6 flex items-center">
						<span class="bg-logitech-blue text-white w-8 h-8 rounded-full flex items-center justify-center mr-3">4</span>
						Preferencias de Contacto
					</h2>
					
					<div class="space-y-6">
						<div>
							<label class="form-label">
								Método de Contacto Preferido
							</label>
							<div class="space-y-3">
								<label class="flex items-center cursor-pointer group">
									<input
										type="radio"
										bind:group={formData.metodo_contacto_preferido}
										value="email"
										class="w-5 h-5 text-logitech-blue focus:ring-logitech-blue"
									/>
									<span class="ml-3 text-gray-700 group-hover:text-logitech-blue transition-colors">Email</span>
								</label>
								<label class="flex items-center cursor-pointer group">
									<input
										type="radio"
										bind:group={formData.metodo_contacto_preferido}
										value="telefono"
										class="w-5 h-5 text-logitech-blue focus:ring-logitech-blue"
									/>
									<span class="ml-3 text-gray-700 group-hover:text-logitech-blue transition-colors">Teléfono</span>
								</label>
								<label class="flex items-center cursor-pointer group">
									<input
										type="radio"
										bind:group={formData.metodo_contacto_preferido}
										value="ambos"
										class="w-5 h-5 text-logitech-blue focus:ring-logitech-blue"
									/>
									<span class="ml-3 text-gray-700 group-hover:text-logitech-blue transition-colors">Ambos</span>
								</label>
							</div>
						</div>

						<div>
							<label for="horario_contacto" class="form-label">
								Horario Preferido para Contacto
							</label>
							<select
								id="horario_contacto"
								bind:value={formData.horario_contacto}
								class="form-input"
							>
								<option value="">Selecciona un horario</option>
								{#each horariosContacto as horario}
									<option value={horario}>{horario}</option>
								{/each}
							</select>
						</div>
					</div>
				</div>

				<!-- Términos y Condiciones -->
				<div class="space-y-4">
					<div>
						<label class="flex items-start cursor-pointer group">
							<input
								type="checkbox"
								bind:checked={formData.acepta_terminos}
								class="w-5 h-5 text-logitech-blue focus:ring-logitech-blue mt-1"
								class:border-red-500={errors.acepta_terminos}
							/>
							<span class="ml-3 text-gray-700 group-hover:text-logitech-blue transition-colors">
								Acepto los <a href="#" class="text-logitech-blue underline hover:no-underline">términos y condiciones</a> y la 
								<a href="#" class="text-logitech-blue underline hover:no-underline">política de privacidad</a> 
								<span class="text-red-500">*</span>
							</span>
						</label>
						{#if errors.acepta_terminos}
							<p class="form-error ml-8">{errors.acepta_terminos}</p>
						{/if}
					</div>

					<div>
						<label class="flex items-start cursor-pointer group">
							<input
								type="checkbox"
								bind:checked={formData.acepta_marketing}
								class="w-5 h-5 text-logitech-blue focus:ring-logitech-blue mt-1"
							/>
							<span class="ml-3 text-gray-700 group-hover:text-logitech-blue transition-colors">
								Deseo recibir información sobre productos, ofertas y novedades de Logitech
							</span>
						</label>
					</div>
				</div>

				<!-- Botones -->
				<div class="flex flex-col sm:flex-row gap-4 pt-6">
					<button
						type="submit"
						disabled={isSubmitting}
						class="btn-primary flex-1 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
					>
						{#if isSubmitting}
							<svg class="animate-spin -ml-1 mr-3 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
								<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
								<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
							</svg>
							Enviando...
						{:else}
							Enviar Formulario
						{/if}
					</button>
					<button
						type="button"
						on:click={() => window.location.reload()}
						class="btn-secondary flex-1"
					>
						Limpiar Formulario
					</button>
				</div>
			</form>
		</div>

		<!-- Footer -->
		<div class="text-center mt-12 text-gray-400">
			<p>© 2024 Logitech. Todos los derechos reservados.</p>
			<p class="mt-2 text-sm">Formulario de soporte técnico - Testing de manejo de formularios</p>
		</div>
	</div>
</div>

<style>
	:global(body) {
		margin: 0;
		padding: 0;
	}
</style>
